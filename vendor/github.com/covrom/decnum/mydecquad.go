package decnum

/*

#include "mydecquad.h"
*/
import "C"

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
	"unsafe"
)

func assert(val bool) {
	if val == false {
		panic("assertion failed")
	}
}

/************************************************************************/
/*                                                                      */
/*                           Quad value and status                      */
/*                                                                      */
/************************************************************************/

// Quad contains a 128bits decimal floating-point value, and a status describing the exceptional conditions that occurred during the generation of the value.
//
// Exceptional conditions can be errors, like 'DivisionByZero', or just informational flags, like 'Inexact'.
//
type Quad C.Quad // array of 16 bytes, + 2 bytes for status field

// Status returns the status field of the Quad.
// It contains error flags, like 'DivisionByZero', and informational flags, like 'Inexact'.
//
func (a Quad) Status() Status {

	return Status(a.status)
}

// ErrorStatus returns the status field of the Quad.
// It contains only error flags.
// Same as a.Status()&ErrorMask
//
func (a Quad) ErrorStatus() Status {

	return Status(a.status) & ErrorMask
}

// Error returns an error if an error flag bit has been set in Quad's status field.
// Many error flag bits can be set in error.
//
// It returns QuadError(a.ErrorStatus()), or nil if no error.
//
func (a Quad) Error() error {
	var errorFlags Status

	errorFlags = Status(a.status) & ErrorMask

	if errorFlags == 0 {
		return nil
	}

	return QuadError(errorFlags)
}

func newError(status Status) QuadError {

	return QuadError(status & ErrorMask)
}

/************************************************************************/
/*                                                                      */
/*                 global constants and variables                       */
/*                                                                      */
/************************************************************************/

const (
	DecquadPmax   = C.DECQUAD_Pmax   // number of digits in coefficient == 34
	DecquadBytes  = C.DECQUAD_Bytes  // size in bytes of decQuad == 16
	DecquadString = C.DECQUAD_String // buffer capacity for C.decQuadToString()
)

/************************************************************************/
/*                                                                      */
/*                              Context                                 */
/*                                                                      */
/************************************************************************/

// Quad contains a value and a status field.
// The status field describes the exceptional conditions that occurred during the calculation.
// During a long calculation series, status bits are set and propagate into the result number.
// This way, you can check the status of a number after a series of calculation, using Error() method.
//
type Status uint16

// These exceptional condition constants are bit flags, power of two.
// They are error flags, or informational flags.
// The only informational flag used by this package is 'Inexact'.
//
const (
	ConversionSyntax    Status = C.DEC_Conversion_syntax    // error flag
	DivisionByZero      Status = C.DEC_Division_by_zero     // error flag
	DivisionImpossible  Status = C.DEC_Division_impossible  // error flag
	DivisionUndefined   Status = C.DEC_Division_undefined   // error flag
	InsufficientStorage Status = C.DEC_Insufficient_storage // error flag
	Inexact             Status = C.DEC_Inexact              // informational flag. It is the only informational flag that can be set by Quad operations.
	InvalidContext      Status = C.DEC_Invalid_context      // error flag
	InvalidOperation    Status = C.DEC_Invalid_operation    // error flag
	Overflow            Status = C.DEC_Overflow             // error flag
	Clamped             Status = C.DEC_Clamped              // informational flag. Quad doesn't use it.
	Rounded             Status = C.DEC_Rounded              // informational flag. Quad doesn't use it.
	Subnormal           Status = C.DEC_Subnormal            // informational flag. Quad doesn't use it.
	Underflow           Status = C.DEC_Underflow            // error flag. E.g. 1e-6000/1e1000

	//LostDigits          Status = C.DEC_Lost_digits        // informational flag. Exists only if DECSUBSET is set, which is not the case by default
)

const ErrorMask Status = C.DEC_Errors // ErrorMask is the bitmask of the error flags, ORed together. After a series of operations, if status & decnum.ErrorMask != 0, an error has occured, e.g. division by 0.

// String representation of a single flag (status with one bit set).
//
func (flag Status) flagString() string {

	if flag == 0 {
		return ""
	}

	switch flag {
	case ConversionSyntax:
		return "ConversionSyntax"
	case DivisionByZero:
		return "DivisionByZero"
	case DivisionImpossible:
		return "DivisionImpossible"
	case DivisionUndefined:
		return "DivisionUndefined"
	case InsufficientStorage:
		return "InsufficientStorage"
	case Inexact:
		return "Inexact"
	case InvalidContext:
		return "InvalidContext"
	case InvalidOperation:
		return "InvalidOperation"
	case Overflow:
		return "Overflow"
	case Clamped:
		return "Clamped"
	case Rounded:
		return "Rounded"
	case Subnormal:
		return "Subnormal"
	case Underflow:
		return "Underflow"
	default:
		return "Unknown status flag"
	}
}

// String representation of a status.
// status can have many flags set.
//
func (status Status) String() string {
	var (
		s    string
		flag Status
	)

	for i := Status(0); i < 32; i++ {
		flag = Status(0x0001 << i)
		if status&flag != 0 {
			if s == "" {
				s = flag.flagString()
			} else {
				s += ";" + flag.flagString()
			}
		}
	}

	return s
}

// error type for this package.
//
type QuadError Status

// Error returns a string describing the error flags.
//
func (e QuadError) Error() string {

	return fmt.Sprintf("decnum: %s", (Status(e) & ErrorMask).String())
}

/************************************************************************/
/*                                                                      */
/*                             rounding mode                            */
/*                                                                      */
/************************************************************************/

type RoundingMode int

// Rounding mode is used if rounding is necessary during an operation.
const (
	RoundCeiling  RoundingMode = C.DEC_ROUND_CEILING   // Round towards +Infinity.
	RoundDown     RoundingMode = C.DEC_ROUND_DOWN      // Round towards 0 (truncation).
	RoundFloor    RoundingMode = C.DEC_ROUND_FLOOR     // Round towards –Infinity.
	RoundHalfDown RoundingMode = C.DEC_ROUND_HALF_DOWN // Round to nearest; if equidistant, round down.
	RoundHalfEven RoundingMode = C.DEC_ROUND_HALF_EVEN // Round to nearest; if equidistant, round so that the final digit is even.
	RoundHalfUp   RoundingMode = C.DEC_ROUND_HALF_UP   // Round to nearest; if equidistant, round up.
	RoundUp       RoundingMode = C.DEC_ROUND_UP        // Round away from 0.
	Round05Up     RoundingMode = C.DEC_ROUND_05UP      // The same as RoundUp, except that rounding up only occurs if the digit to be rounded up is 0 or 5 and after Overflow the result is the same as for RoundDown.
	RoundDefault  RoundingMode = RoundHalfEven         // The same as RoundHalfEven.
)

func (rounding RoundingMode) String() string {

	switch rounding {
	case RoundCeiling:
		return "RoundCeiling"
	case RoundDown:
		return "RoundDown"
	case RoundFloor:
		return "RoundFloor"
	case RoundHalfDown:
		return "RoundHalfDown"
	case RoundHalfEven:
		return "RoundHalfEven"
	case RoundHalfUp:
		return "RoundHalfUp"
	case RoundUp:
		return "RoundUp"
	case Round05Up:
		return "Round05Up"
	default:
		return "Unknown rounding mode"
	}
}

/************************************************************************/
/*                                                                      */
/*                       init and version functions                     */
/*                                                                      */
/************************************************************************/

var (
	decNumberVersion string = C.GoString(C.decQuadVersion()) // version of the original C decNumber package

	decNumberMacros string = fmt.Sprintf("decQuad module: DECDPUN %d, DECSUBSET %d, DECEXTFLAG %d. Constants DECQUAD_Pmax %d, DECQUAD_String %d DECQUAD_Bytes %d.",
		C.DECDPUN, C.DECSUBSET, C.DECEXTFLAG, C.DECQUAD_Pmax, C.DECQUAD_String, C.DECQUAD_Bytes) // macros defined by the C decNumber module
)

func init() {
	C.mdq_init()

	if DecquadBytes != 16 { // 16 bytes == 128 bits
		panic("DECQUAD_Bytes != 16")
	}

	assert(C.DECSUBSET == 0) // because else, we should define LostDigits as status flag

	assert(poolBuffCapacity > DecquadPmax)
	assert(poolBuffCapacity > DecquadString)
}

// DecNumberVersion returns the version of the original C decNumber package.
//
// This function is only useful if you want information about the original C decNumber package.
//
func DecNumberVersion() string {
	return decNumberVersion
}

// DecNumberMacros returns the values of macros defined in the original C decNumber package.
//
// This function is only useful if you want information about the original C decNumber package.
//
func DecNumberMacros() string {
	return decNumberMacros
}

// g_nan, g_zero and g_one are private variable, because else, a user of the package can change their value by doing decnum.G_ZERO = ...

var (
	g_nan  Quad = nan_for_varinit()     // a constant Quad with value NaN. It runs BEFORE init().
	g_zero Quad = zero_for_varinit()    // a constant Quad with value 0.   It runs BEFORE init().
	g_one  Quad = quad_for_varinit("1") // a constant Quad with value 1.   It runs BEFORE init().
)

// used only to initialize the global variable g_nan.
//
// So, it runs BEFORE init().
//
func nan_for_varinit() (r Quad) {
	var val C.decQuad

	val = C.mdq_nan()

	return Quad{val: val, status: 0}
}

// used only to initialize the global variable g_zero.
//
// So, it runs BEFORE init().
//
func zero_for_varinit() (r Quad) {
	var val C.decQuad

	val = C.mdq_zero()

	return Quad{val: val, status: 0}
}

// used only to initialize some global variables, like g_one.
//
// So, it runs BEFORE init().
//
func quad_for_varinit(s string) (r Quad) {
	var val Quad
	var err error

	if val, err = FromString(s); err != nil {
		panic("decnum: initialization error in quad_for_varinit() " + err.Error())
	}

	return val
}

/************************************************************************/
/*                                                                      */
/*                      arithmetic operations                           */
/*                                                                      */
/************************************************************************/

type CmpFlag uint32 // result of Compare

const (
	CmpLess    CmpFlag = C.CMP_LESS    // 1
	CmpEqual   CmpFlag = C.CMP_EQUAL   // 2
	CmpGreater CmpFlag = C.CMP_GREATER // 4
	CmpNaN     CmpFlag = C.CMP_NAN     // 8
)

func (cmp CmpFlag) String() string {

	switch cmp {
	case CmpLess:
		return "CmpLess"
	case CmpEqual:
		return "CmpEqual"
	case CmpGreater:
		return "CmpGreater"
	case CmpNaN:
		return "CmpNaN"
	default:
		return "Unknown CmpFlag"
	}
}

// GetExponent can return these special values for NaN, sNaN, Infinity.
const (
	ExpNaN          = C.DECFLOAT_NaN
	ExpSignalingNaN = C.DECFLOAT_sNaN
	ExpInf          = C.DECFLOAT_Inf
	ExpMinSpecial   = C.DECFLOAT_MinSp // minimum special value. Special values are all >= ExpMinSpecial
)

// Zero returns 0 Quad value.
//
//     r = Zero()  // assign 0 to the Quad r
//
func Zero() (r Quad) {

	return g_zero
}

// One returns 1 Quad value.
//
//     r = One()  // assign 1 to the Quad r
//
func One() (r Quad) {

	return g_one
}

// NaN returns NaN Quad value.
//
//     r = NaN()  // assign NaN to the Quad r
//
func NaN() (r Quad) {

	return g_nan
}

// Copy returns a copy of a.
//
// But it is easier to just use '=' :
//
//        a = r
//
func Copy(a Quad) Quad {

	return a
}

// ClearStatus returns a copy of a, whith status field cleared.
//
func (a Quad) ClearStatus() Quad {

	a.status = 0
	return a
}

// SetStatusFlags returns a copy of a, setting the status flags specified by argument.
// Status of result is a.status|statusflags.
//
// Normally, only library and test modules use this function.
//
func (a Quad) SetStatusFlags(statusflags Status) Quad {

	a.status |= C.uint16_t(statusflags)

	return a
}

// ClearStatusFlags returns a copy of a, clearing the status flags specified by argument.
// Status of result is a.status&^statusflags.
//
// Normally, only library and test modules use this function.
//
func (a Quad) ClearStatusFlags(statusflags Status) Quad {

	a.status &^= C.uint16_t(statusflags)

	return a
}

// Neg returns -a.
//
func (a Quad) Neg() Quad {

	return Quad(C.mdq_minus(C.struct_Quad(a)))
}

// Add returns a + b.
//
func (a Quad) Add(b Quad) Quad {

	return Quad(C.mdq_add(C.struct_Quad(a), C.struct_Quad(b)))
}

// Sub returns a - b.
//
func (a Quad) Sub(b Quad) Quad {

	return Quad(C.mdq_subtract(C.struct_Quad(a), C.struct_Quad(b)))
}

// Mul returns a * b.
//
func (a Quad) Mul(b Quad) Quad {

	return Quad(C.mdq_multiply(C.struct_Quad(a), C.struct_Quad(b)))
}

// Div returns a/b.
//
func (a Quad) Div(b Quad) Quad {

	return Quad(C.mdq_divide(C.struct_Quad(a), C.struct_Quad(b)))
}

// DivInt returns the integral part of a/b.
//
func (a Quad) DivInt(b Quad) Quad {

	return Quad(C.mdq_divide_integer(C.struct_Quad(a), C.struct_Quad(b)))
}

// Mod returns the modulo of a and b.
//
func (a Quad) Mod(b Quad) Quad {

	return Quad(C.mdq_remainder(C.struct_Quad(a), C.struct_Quad(b)))
}

// Max returns the larger of a and b.
// If either a or b is NaN then the other argument is the result.
//
func Max(a Quad, b Quad) Quad {

	return Quad(C.mdq_max(C.struct_Quad(a), C.struct_Quad(b)))
}

// Min returns the smaller of a and b.
// If either a or b is NaN then the other argument is the result.
//
func Min(a Quad, b Quad) Quad {

	return Quad(C.mdq_min(C.struct_Quad(a), C.struct_Quad(b)))
}

// ToIntegral returns the value of a rounded to an integral value.
//
//      The representation of a number is:
//
//           (-1)^sign  coefficient * 10^exponent
//           where coefficient is an integer storing 34 digits.
//
//       - If exponent < 0, the least significant digits are discarded, so that new exponent becomes 0.
//             Internally, it calls Quantize(a, 1E0) with specified rounding.
//       - If exponent >= 0, the number remains unchanged.
//
//         E.g.     12.345678e2    is     12345678E-4     -->   1235E0
//                  123e5          is     123E5        remains   123E5
//
// See also Round, RoundMode and Truncate methods, which are more convenient to use.
//
func (a Quad) ToIntegral(rounding RoundingMode) Quad {

	return Quad(C.mdq_to_integral(C.struct_Quad(a), C.int(rounding)))
}

// Quantize rounds a to the same pattern as b.
// b is just a model, its sign and coefficient value are ignored. Only its exponent is used.
// The result is the value of a, but with the same exponent as the pattern b.
//
//      The representation of a number is:
//
//           (-1)^sign  coefficient * 10^exponent
//           where coefficient is an integer storing 34 digits.
//
// Examples (with RoundHalfEven rounding mode):
//    quantization of 134.6454 with    0.00001    is   134.64540
//                    134.6454 with    0.00000    is   134.64540     the value of b has no importance
//                    134.6454 with 1234.56789    is   134.64540     the value of b has no importance
//                    134.6454 with 0.0001        is   134.6454
//                    134.6454 with 0.01          is   134.65
//                    134.6454 with 1             is   135
//                    134.6454 with 1000000000    is   135           the value of b has no importance
//                    134.6454 with 1E+2          is   1E+2
//
//		        123e32 with 1             sets Invalid_operation error flag in status
//		        123e32 with 1E1           is   1230000000000000000000000000000000E1
//		        123e32 with 10            sets Invalid_operation error flag in status
//
// See also Round, RoundMode and Truncate methods, which are more useful methods.
//
func (a Quad) Quantize(b Quad, rounding RoundingMode) Quad {

	return Quad(C.mdq_quantize(C.struct_Quad(a), C.struct_Quad(b), C.int(rounding)))
}

// Abs returns the absolute value of a.
//
func (a Quad) Abs() Quad {

	return Quad(C.mdq_abs(C.struct_Quad(a)))
}

/************************************************************************/
/*                                                                      */
/*                            IsFinite, etc                             */
/*                                                                      */
/************************************************************************/

// IsFinite returns true if a is not Infinite, nor Nan.
//
// The status field of a is not checked.
// If you need to check the status of a, you can call a.Error().
//
func (a Quad) IsFinite() bool {

	if C.mdq_is_finite(a.val) != 0 {
		return true
	}

	return false
}

/* IsInteger is discarded.
   I keep the code here just in case my opinion changes, but I think method is error-prone for the user.
   E.g. 1e3 returns false. I am not sure it is what the user intuitively expects.

// IsInteger returns true if a is finite and has exponent=0.
//
//      The number representation is:
//
//           (-1)^sign  coefficient * 10^exponent
//           where coefficient is an integer storing 34 digits.
//
//      If the number in the above representation has exponent=0, then IsInteger returns true.
//
//      0              0E+0        returns true
//      1              1E+0        returns true
//      12.34e2     1234E+0        returns true
//
//      0.0000         0E-4        returns false
//      1.0000     10000E-4        returns false
//     -12.34e5    -1234E+3        returns false
//      1e3            1E+3        returns false
//
func (a Quad) IsInteger() bool {

	if C.mdq_is_integer(a.val) != 0 {
		return true
	}

	return false
}
*/

// IsInfinite returns true if a is Infinite.
//
// The status field of a is not checked.
// If you need to check the status of a, you can call a.Error().
//
func (a Quad) IsInfinite() bool {

	if C.mdq_is_infinite(a.val) != 0 {
		return true
	}

	return false
}

// IsNaN returns true if a is Nan.
//
// The status field of a is not checked.
// If you need to check the status of a, you can call a.Error().
//
func (a Quad) IsNaN() bool {

	if C.mdq_is_nan(a.val) != 0 {
		return true
	}

	return false
}

// IsPositive returns true if a > 0 and not Nan.
//
// The status field of a is not checked.
// If you need to check the status of a, you can call a.Error().
//
func (a Quad) IsPositive() bool {

	if C.mdq_is_positive(a.val) != 0 {
		return true
	}

	return false
}

// IsZero returns true if a == 0.
//
// The status field of a is not checked.
// If you need to check the status of a, you can call a.Error().
//
func (a Quad) IsZero() bool {

	if C.mdq_is_zero(a.val) != 0 {
		return true
	}

	return false
}

// IsNegative returns true if a < 0 and not NaN.
//
// The status field of a is not checked.
// If you need to check the status of a, you can call a.Error().
//
func (a Quad) IsNegative() bool {

	if C.mdq_is_negative(a.val) != 0 {
		return true
	}

	return false
}

// GetExponent returns the exponent of a.
//
//      The representation of a number is:
//
//           (-1)^sign  coefficient * 10^exponent
//           where coefficient is an integer storing 34 digits.
//
// This function returns the exponent.
// It can returns special values such as ExpNaN, ExpSignalingNaN or ExpInf if a is NaN, sNaN or Infinity.
//
func (a Quad) GetExponent() int32 {

	return int32(C.mdq_get_exponent(a.val))
}

/************************************************************************/
/*                                                                      */
/*                            comparison                                */
/*                                                                      */
/************************************************************************/

// Greater is true if a > b.
//
// The status fields of a and b are not checked.
// If you need to check them, you can call a.Error() and b.Error().
//
func (a Quad) Greater(b Quad) bool {
	var result C.uint32_t

	result = C.mdq_compare(C.struct_Quad(a), C.struct_Quad(b))

	if CmpFlag(result)&CmpGreater != 0 {
		return true
	}

	return false
}

// GreaterEqual is true if a >= b.
//
// The status fields of a and b are not checked.
// If you need to check them, you can call a.Error() and b.Error().
//
func (a Quad) GreaterEqual(b Quad) bool {
	var result C.uint32_t

	result = C.mdq_compare(C.struct_Quad(a), C.struct_Quad(b))

	if CmpFlag(result)&(CmpGreater|CmpEqual) != 0 {
		return true
	}

	return false
}

// Equal is true if a == b.
//
// The status fields of a and b are not checked.
// If you need to check them, you can call a.Error() and b.Error().
//
func (a Quad) Equal(b Quad) bool {
	var result C.uint32_t

	result = C.mdq_compare(C.struct_Quad(a), C.struct_Quad(b))

	if CmpFlag(result)&CmpEqual != 0 {
		return true
	}

	return false
}

// LessEqual is true if a <= b.
//
// The status fields of a and b are not checked.
// If you need to check them, you can call a.Error() and b.Error().
//
func (a Quad) LessEqual(b Quad) bool {
	var result C.uint32_t

	result = C.mdq_compare(C.struct_Quad(a), C.struct_Quad(b))

	if CmpFlag(result)&(CmpLess|CmpEqual) != 0 {
		return true
	}

	return false
}

// Less is true if a < b.
//
// The status fields of a and b are not checked.
// If you need to check them, you can call a.Error() and b.Error().
//
func (a Quad) Less(b Quad) bool {
	var result C.uint32_t

	result = C.mdq_compare(C.struct_Quad(a), C.struct_Quad(b))

	if CmpFlag(result)&CmpLess != 0 {
		return true
	}

	return false
}

/************************************************************************/
/*                                                                      */
/*                   conversion from string and numbers                 */
/*                                                                      */
/************************************************************************/

// FromString returns a Quad from a string.
//
// Special values "NaN" (also "qNaN"), "sNaN", "NaN123" (NaN with payload), "sNaN123" (sNaN with payload), "Infinity" (or "Inf", "+Inf"), "-Infinity" ( or "-Inf") are accepted.
//
//      Infinity and -Infinity, or Inf and -Inf, represent a value infinitely large.
//
//      NaN or qNaN, which means "Not a Number", represents an undefined result, when an arithmetic operation has failed. E.g. FromString("hello")
//                   NaN propagates to all subsequent operations, because if NaN is passed as argument, the result, will be NaN.
//                   These NaN are called "quiet NaN", because they don't set exceptional condition flag in status when passed as argument to an operation.
//
//      sNaN, or "signaling NaN", are created by FromString("sNaN"). When passed as argument to an operation, the result will be NaN, like with quiet NaN.
//                   But they will set (==signal) an exceptional condition flag in status, "Invalid_operation".
//                   Signaling NaN propagate to subsequent operation as ordinary NaN (quiet NaN), and not as "signaling NaN".
//
// Note that both NaN and sNaN can take an integer payload, e.g. NaN123, created by FromString("NaN123"), and it is up to you to give it a significance.
// sNaN and payload are not used often, and most probably, you won't use them.
//
// This function returns result.Error() as a convenience.
//
func FromString(s string) (result Quad, err error) {
	var cs *C.char

	s = strings.TrimSpace(s)

	cs = C.CString(s)
	defer C.free(unsafe.Pointer(cs))

	result = Quad(C.mdq_from_string(cs))

	return result, result.Error()
}

// FromInt32 returns a Quad from a int32 value.
//
// No error occurs.
//
func FromInt32(value int32) Quad {

	return Quad(C.mdq_from_int32(C.int32_t(value)))
}

// FromInt64 returns a Quad from a int64 value.
//
// No error occurs.
//
// Note that FromInt64 is slower than FromInt32, because the underlying C decNumber package has no function that converts directly from int64.
// So, int64 is first converted to string, and then to Quad.
//
func FromInt64(value int64) Quad {

	return Quad(C.mdq_from_int64(C.int64_t(value)))
}

func FromFloat(value float64) Quad {
	q, err := FromString(strconv.FormatFloat(value, 'f', -1, 64))
	if err != nil {
		panic(err)
	}
	return q
}

/************************************************************************/
/*                                                                      */
/*                      conversion to string                            */
/*                                                                      */
/************************************************************************/

const poolBuffCapacity = 50 // capacity of []byte buffer generated by the pool of buffers

// pool is a pool of byte slice, used by AppendQuad and String.
//
// note:
//    DecquadString      = 43         sign, 34 digits, decimal point, E+xxxx, terminal \0   gives 43
//    DecquadPmax        = 34
//    poolBuffCapacity   = 50         just to be sure, it is largely enough
//
// The pool must return []byte with capacity being at least the largest of DecquadString and DecquadPmax. We prefer a capacity of poolBuffCapacity to be sure.
//
var pool = sync.Pool{
	New: func() interface{} {
		//fmt.Println("---   POOL")
		return make([]byte, poolBuffCapacity) // poolBuffCapacity is larger than DecquadString and DecquadPmax. This size is ok for AppendQuad and String methods.
	},
}

// QuadToString returns the string representation of a Quad number.
// It calls the C function QuadToString of the original decNumber package.
//
//       This function uses exponential notation quite often.
//       E.g. 0.0000001 returns "1E-7", which is often not what we want.
//
//       It is better to use the method AppendQuad() or String(), which don't use exponential notation for a wider range.
//       AppendQuad() and String() write a number without exp notation if it can be displayed with at most 34 digits, and an optional fractional point.
//
// The status field of a is not checked.
// If you need to check the status of a, you can call a.Error().
//
func (a Quad) QuadToString() string {
	var (
		retStr   C.Ret_str
		strSlice []byte // capacity must be exactly DecquadString
		s        string
	)

	retStr = C.mdq_QuadToString(a.val) // may use exponent notation

	strSlice = pool.Get().([]byte)[:DecquadString]
	defer pool.Put(strSlice)

	for i := 0; i < int(retStr.length); i++ {
		strSlice[i] = byte(retStr.s[i])
	}

	s = string(strSlice[:retStr.length])

	return s
}

// AppendQuad appends string representation of Quad into byte slice.
// AppendQuad and String are best to display Quad, as exponent notation is used less often than with QuadToString.
//
//       AppendQuad() writes a number without exp notation if it can be displayed with at most 34 digits, and an optional fractional point.
//       Else, falls back on QuadToString(), which will use exponential notation.
//
// See also method String(), which calls AppendQuad internally.
//
// The status field of a is not checked.
// If you need to check the status of a, you can call a.Error().
//
func AppendQuad(dst []byte, a Quad) []byte {
	var (
		retStr   C.Ret_str
		strSlice []byte // length must be exactly DecquadString

		ret             C.Ret_BCD
		d               byte
		skipLeadingZero bool = true
		inf_nan         uint32
		exp             int32
		sign            uint32
		BCDslice        []byte // length must be exactly DecquadPmax

		buff [DecquadString]byte // enough for      sign    optional "0."    34 digits
	)

	// fill BCD array

	ret = C.mdq_to_BCD(a.val) // sign will be 1 for negative and non-zero number, else, 0. If Inf or Nan, returns an error.

	BCDslice = pool.Get().([]byte)[:DecquadPmax]
	defer pool.Put(BCDslice)

	for i := 0; i < DecquadPmax; i++ {
		BCDslice[i] = byte(ret.BCD[i])
	}
	inf_nan = uint32(ret.inf_nan)
	exp = int32(ret.exp)
	sign = uint32(ret.sign)

	// if Quad value is not in 34 digits range, or Inf or Nan, we want our function to output the number, or Infinity, or NaN. Falls back on QuadToString.

	if exp > 0 || exp < -DecquadPmax || inf_nan != 0 {
		retStr = C.mdq_QuadToString(a.val) // may use exponent notation

		strSlice = pool.Get().([]byte)[:DecquadString]
		defer pool.Put(strSlice)

		for i := 0; i < int(retStr.length); i++ {
			strSlice[i] = byte(retStr.s[i])
		}

		dst = append(dst, strSlice[:retStr.length]...) // write buff into destination and return

		return dst
	}

	// write string. Here, the number is not Inf nor Nan.

	i := 0

	integralPartLength := len(BCDslice) + int(exp) // here, exp is [-DecquadPmax ... 0]

	BCDintegralPart := BCDslice[:integralPartLength]
	BCDfractionalPart := BCDslice[integralPartLength:]

	for _, d = range BCDintegralPart { // ==== write integral part ====
		if skipLeadingZero && d == 0 {
			continue
		} else {
			skipLeadingZero = false
		}
		buff[i] = '0' + d
		i++
	}

	if i == 0 { // write '0' if no digit written for integral part
		buff[i] = '0'
		i++
	}

	if sign != 0 {
		dst = append(dst, '-') // write '-' sign if any into destination
	}

	dst = append(dst, buff[:i]...) // write integral part into destination

	if exp == 0 { // if no fractional part, just return
		return dst
	}

	dst = append(dst, '.') // ==== write fractional part ====

	i = 0
	for _, d = range BCDfractionalPart {
		buff[i] = '0' + d
		i++
	}

	dst = append(dst, buff[:i]...) // write fractional part into destination

	return dst
}

// String is the preferred way to display a decQuad number.
// It calls AppendQuad internally.
//
// The status field of a is not checked.
// If you need to check the status of a, you can call a.Error().
//
func (a Quad) String() string {
	var buffer []byte

	buffer = pool.Get().([]byte)[:0] // capacity is enough to receive result of C.mdq_QuadToString(), and also big enough to receive [sign] + [DecquadPmax digits] + [fractional dot]
	defer pool.Put(buffer)

	ss := AppendQuad(buffer[:0], a)

	return string(ss)
}

/************************************************************************/
/*                                                                      */
/*                      conversion to number                            */
/*                                                                      */
/************************************************************************/

// ToInt32 returns the int32 value from a.
//
// The status field of a is not checked.
// If you need to check the status of a, you can call a.Error().
//
func (a Quad) ToInt32(rounding RoundingMode) (int32, error) {
	var result C.Ret_int32_t

	result = C.mdq_to_int32(C.struct_Quad(a), C.int(rounding))

	if Status(result.status)&ErrorMask != 0 {
		return 0, newError(Status(result.status))
	}

	return int32(result.val), nil
}

// ToInt64 returns the int64 value from a.
// The rounding passed as argument is used, instead of the rounding mode of context which is ignored.
//
// The status field of a is not checked.
// If you need to check the status of a, you can call a.Error().
//
// Note that ToInt64 is slower than ToInt32, because the underlying C decNumber package has no function that converts directly to int64.
// So, the number is first converted to string, and then to int64.
//
func (a Quad) ToInt64(rounding RoundingMode) (int64, error) {
	var result C.Ret_int64_t

	result = C.mdq_to_int64(C.struct_Quad(a), C.int(rounding))

	if Status(result.status)&ErrorMask != 0 {
		return 0, newError(Status(result.status))
	}

	return int64(result.val), nil
}

// ToFloat64 returns the float64 value from a.
//
// The status field of a is not checked.
// If you need to check the status of a, you can call a.Error().
//
func (a Quad) ToFloat64() (float64, error) {
	var (
		err error
		val float64
	)

	if a.IsNaN() { // because strconv.ParseFloat doesn't parse signaling sNaN
		return math.NaN(), nil
	}

	if val, err = strconv.ParseFloat(a.String(), 64); err != nil {
		return math.NaN(), QuadError(InvalidOperation)
	}

	return val, nil
}

// Bytes returns the internal byte representation of the value field of the Quad.
// It is not useful, except for educational purpose.
//
func (a Quad) Bytes() (res [DecquadBytes]byte) {

	for i, b := range a.val {
		res[i] = byte(b)
	}

	return res
}

/************************************************************************/
/*                                                                      */
/*                      rounding and truncating                         */
/*                                                                      */
/************************************************************************/

// RoundWithMode rounds (or truncate) 'a', with the mode passed as argument.
// You must pass a constant RoundCeiling, RoundHalfEven, etc as argument.
//
//  n must be in the range [-35...34]. Else, Invalid Operation flag is set, and NaN is returned.
//
func (a Quad) RoundWithMode(n int32, rounding RoundingMode) Quad {

	return Quad(C.mdq_roundM(C.struct_Quad(a), C.int32_t(n), C.int(rounding)))
}

// Round rounds (or truncate) 'a', with RoundHalfEven mode.
//
//  n must be in the range [-35...34]. Else, Invalid Operation flag is set, and NaN is returned.
//
func (a Quad) Round(n int32) Quad {

	return Quad(C.mdq_roundM(C.struct_Quad(a), C.int32_t(n), C.int(RoundHalfEven)))
}

// Truncate truncates 'a'.
// It is like rounding with RoundDown.
//
//  n must be in the range [-35...34]. Else, Invalid Operation flag is set, and NaN is returned.
//
func (a Quad) Truncate(n int32) Quad {

	return Quad(C.mdq_roundM(C.struct_Quad(a), C.int32_t(n), C.int(RoundDown)))
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (a *Quad) UnmarshalJSON(bytes []byte) error {
	var str string
	if len(bytes) != 0 && bytes[0] == '"' {
		str = string(bytes[1 : len(bytes)-1])
	} else {
		str = string(bytes)
	}

	d, err := FromString(str)
	*a = d
	if err != nil {
		return fmt.Errorf("Error decoding string '%s': %s", str, err)
	}
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (a Quad) MarshalJSON() ([]byte, error) {
	str := "\"" + a.String() + "\""
	return []byte(str), nil
}
