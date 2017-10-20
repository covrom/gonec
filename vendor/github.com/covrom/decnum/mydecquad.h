#ifndef MYDECQUAD_H
#define MYDECQUAD_H

#include <errno.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <assert.h>
#include "decQuad.h"      // this header includes "decContext.h"
#include "decimal128.h"   // interface to decNumber, used for decNumberPower(). Also for definition of DECDPUN.


#define MDQ_INFINITE    1     // result is Inf or -Inf
#define MDQ_NAN         2     // result is Nan


#define CMP_LESS        1
#define CMP_EQUAL       2
#define CMP_GREATER     4
#define CMP_NAN         8


void mdq_init(void);


// Quad is the decimal number structure.
// It contains a 128bits value, and a status field as meta-data to the value.
// The status bits are set by all exceptional conditions encountered during the generation of the value, by a function or operator.
//
// ***** in decNumber library, status field is defined as 32 bits, but status flags constants only need 16 bits (see decContext.h). *****
//
typedef struct Quad {
    decQuad     val;
    uint16_t    status; // see comment above
} Quad;


// struct used to pass BCD string from C to Go, by value.
//
typedef struct Ret_BCD {
  uint32_t   inf_nan;
  uint8_t    BCD[DECQUAD_Pmax];
  int32_t    exp;
  uint32_t   sign;
} Ret_BCD;

// struct used to pass string from C to Go, by value.
//
typedef struct Ret_str {
  char       s[DECQUAD_String];
  size_t     length;
} Ret_str;

// struct used to pass int32 from C to Go, by value.
//
typedef struct Ret_int32_t {
  int32_t     val;
  uint16_t    status;
} Ret_int32_t;

// struct used to pass int64 from C to Go, by value.
//
typedef struct Ret_int64_t {
  int64_t     val;
  uint16_t    status;
} Ret_int64_t;


decQuad       mdq_zero();
decQuad       mdq_nan();


Quad          mdq_minus(Quad a);
Quad          mdq_add(Quad a, Quad b);
Quad          mdq_subtract(Quad a, Quad b);
Quad          mdq_multiply(Quad a, Quad b);
Quad          mdq_divide(Quad a, Quad b);
Quad          mdq_divide_integer(Quad a, Quad b);
Quad          mdq_remainder(Quad a, Quad b);
Quad          mdq_max(Quad a, Quad b);
Quad          mdq_min(Quad a, Quad b);
Quad          mdq_to_integral(Quad a, int round);
Quad          mdq_quantize(Quad a, Quad b, int round);
Quad          mdq_abs(Quad a);

uint32_t      mdq_is_finite(decQuad a);
uint32_t      mdq_is_integer(decQuad a);
uint32_t      mdq_is_infinite(decQuad a);
uint32_t      mdq_is_nan(decQuad a);
uint32_t      mdq_is_positive(decQuad a);
uint32_t      mdq_is_zero(decQuad a);
uint32_t      mdq_is_negative(decQuad a);
int32_t       mdq_get_exponent(decQuad a);

uint32_t      mdq_compare(Quad a, Quad b);

Quad          mdq_from_string(char *s);
Quad          mdq_from_int32(int32_t value);
Quad          mdq_from_int64(int64_t value);

Ret_str       mdq_QuadToString(decQuad a);
Ret_BCD       mdq_to_BCD(decQuad a);
Ret_int32_t   mdq_to_int32(Quad a, int round);
Ret_int64_t   mdq_to_int64(Quad a, int round);

Quad          mdq_roundM(Quad a, int32_t n, int round);


#endif



