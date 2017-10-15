package core

import (
	"runtime"
	"strconv"
	"sync"
	"time"

	"github.com/covrom/gonec/consulapi"
)

// VMMainServiceBus главный менеджер сервисов
var VMMainServiceBus = NewServiceBus()

// VMServiceHeader заголовок сервиса, определяющий его адресуемость
type VMServiceHeader struct {
	ID       string
	Path     string // используется в URL сервисов
	Port     string // число 1-65535
	Name     string // имя
	Tags     []string
	External string //= "consul" регистрирует сервис в Consul
}

func NewServiceBus() *VMServiceBus {
	v := &VMServiceBus{
		Name:     "Главный менеджер сервисов Гонец",
		services: make(map[string]VMServicer),
		done:     make(chan bool),
	}
	return v
}

// VMServiceBus внутренний менеджер сервисов
// работает как роутер: открывает порты, обрабатывает пути, вызывает обработчики сервисов
// взаимодействует с внешними service discovery, если требуется
type VMServiceBus struct {
	sync.RWMutex
	wg     sync.WaitGroup
	runned bool

	Name     string
	services map[string]VMServicer //ключ это Id из VMServiceHeader
	done     chan bool
}

func (x *VMServiceBus) vmval() {}

func (x *VMServiceBus) Interface() interface{} {
	return x
}

func (x *VMServiceBus) String() string {
	return x.Name
}

func (x *VMServiceBus) Stop() {
	if x.runned {
		x.done <- true
	}
}

func (x *VMServiceBus) GetService(id string) (VMServicer, bool) {
	x.RLock()
	defer x.RUnlock()

	v, ok := x.services[id]
	return v, ok
}

// Run запускает сервис менеджера сервисов
// он каждую секунду стартует зарегистрированные сервисы, если они не прошли HealthCheck
// проверяет остановку менеджера, и если она произошла - останавливает все живые сервисы
func (x *VMServiceBus) Run() {
	if x.runned {
		return
	}
	x.wg.Add(1)
	x.runned = true

	go func(bus *VMServiceBus) {
		defer bus.wg.Done()
		defer func() {
			bus.runned = false
		}()

		for {
			bus.RLock()
			if len(bus.services) == 0 {
				// не осталось ни одного сервиса - выходим
				bus.RUnlock()
				return
			}
			for _, svc := range bus.services {
				if svc.HealthCheck() != nil {
					svc.Start()
				}
			}
			bus.RUnlock()
			select {
			case <-bus.done:
				// останавливаем все живые сервисы
				bus.RLock()
				for _, svc := range bus.services {
					if svc.HealthCheck() == nil {
						svc.Stop()
					}
				}
				bus.RUnlock()
				return

			case <-time.After(time.Second):
				break
			}
			runtime.Gosched()
		}
	}(x)
}

// Register регистрирует сервис в менеджере, дальше с ним можно работать по ID
func (x *VMServiceBus) Register(svc VMServicer) error {
	x.Lock()
	defer x.Unlock()

	hdr := svc.Header()
	if _, ok := x.services[hdr.ID]; ok {
		return VMErrorServiceAlreadyRegistered
	}
	x.services[hdr.ID] = svc

	if hdr.External == "consul" {
		cli, err := consulapi.NewClient(consulapi.DefaultConfig())
		if err != nil {
			return err
		}
		p, err := strconv.Atoi(hdr.Port)
		if err != nil {
			return err
		}

		cat := cli.Catalog()
		reg := &consulapi.CatalogRegistration{
			ID: hdr.ID,
			Service: &consulapi.AgentService{
				ID: hdr.Path,
				Service: hdr.Name,
				Port:    p,
			},
		}
		_, err = cat.Register(reg, nil)
		if err != nil {
			return err
		}
	}

	return nil
}

// Register апдейтит сервис в менеджере, останавливая и удаляя старую версию
func (x *VMServiceBus) UpdateRegister(svc VMServicer) error {
	x.Lock()
	defer x.Unlock()

	id := svc.Header().ID
	if v, ok := x.services[id]; ok {
		x.Deregister(v)
	}
	x.services[id] = svc
	return nil
}

// Register останавливает (если жив) и удаляет сервис из менеджера
func (x *VMServiceBus) Deregister(svc VMServicer) error {
	x.Lock()
	defer x.Unlock()

	hdr := svc.Header()
	v, ok := x.services[hdr.ID]
	if ok {
		if v.HealthCheck() == nil {
			v.Stop()
		}
	}
	delete(x.services, hdr.ID)

	if hdr.External == "consul" {
		cli, err := consulapi.NewClient(consulapi.DefaultConfig())
		if err != nil {
			return err
		}
		cat := cli.Catalog()
		reg := &consulapi.CatalogDeregistration{
			ServiceID: hdr.ID,
		}
		_, err = cat.Deregister(reg, nil)
		if err != nil {
			return err
		}
	}

	return nil
}

// WaitForAll ожидает завершения работы всех сервисов
func (x *VMServiceBus) WaitForAll() {
	if x.runned {
		x.wg.Wait()
	}
}