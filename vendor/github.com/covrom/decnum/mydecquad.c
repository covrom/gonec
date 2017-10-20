#include "mydecquad.h"


/************************************************************************/
/*                 global constant for Round and Truncate               */
/************************************************************************/

/* decNumber constants for decQuad rounding.

   It contains
      - 1e0
      - 1e-1
      - 1e-2
      - ...
      - 1e-DECQUAD_Pmax        (1e-34)
*/
static decQuad G_DECQUAD_QUANTIZER[DECQUAD_Pmax+1];  // 0...34


/* decNumber constants for decQuad rounding.

   It contains
      - 1e0
      - 1e1
      - 1e2
      - ...
      - 1eDECQUAD_Pmax         (1e34)
      - 1e(DECQUAD_Pmax+1)     (1e35)

    NOTE: the max index is 35, because rounding      9234567890123456789012345678901234
                                           with     10000000000000000000000000000000000    (1e34)
                                           gives    10000000000000000000000000000000000

                                   and rounding      9234567890123456789012345678901234
                                           with    100000000000000000000000000000000000    (1e35)
                                           gives                                      0

                                   So, we must allow rounding functions to use an integral part quantizer of 1e35.
*/
static decQuad G_DECQUAD_INTEGRAL_PART_QUANTIZER[DECQUAD_Pmax+2];  // 0...35


/************************************************************************/
/*                          init and context                            */
/************************************************************************/

static decQuad static_one;  // contains 1, only used by mdq_to_int64


/* initialize the global constants used by this library.

   It is called by Go in init() function.

   Exit(1) if an error occurs.
*/
void mdq_init(void) {

  decContext   set;
  const char  *s;
  int          i;


  //----- check DECLITEND -----

  if ( decContextTestEndian(1) ) {  // if argument is 0, a warning message is displayed (using printf) if DECLITEND is set incorrectly. If 1, no message is displayed. Returns 0 if correct.
      fprintf(stderr, "INITIALIZATION mydecquad.c:mdq_init() FAILED: decnum: decContextTestEndian() failed. Change DECLITEND constant (see \"The decNumber Library\")");
      exit(1);
  }

  assert( DECQUAD_Pmax == 34 );             // we have 34 digits max precision (number of significant digits).
  assert( DECQUAD_String > DECQUAD_Pmax );  // because Go function quad.AppendQuad() requires it


  //----- put 1 in static_one -----

  decQuadFromInt32(&static_one, 1); // IMPORTANT: this means that mdq_to_int64 can only be called after Go init() has been run, as it uses static_one. Method ToInt32() cannot be called to initialize Go global variables.


  //----- fill decContext -----

  decContextDefault(&set, DEC_INIT_DECQUAD);

  if ( decContextGetRounding(&set) != DEC_ROUND_HALF_EVEN ) {
      fprintf(stderr, "INITIALIZATION mydecquad.c:mdq_init() FAILED: decnum: decContextGetRounding(&set) != DEC_ROUND_HALF_EVEN");
      exit(1);
  }


  //----- fill G_DECQUAD_QUANTIZER[] -----

  decQuadFromInt32(&G_DECQUAD_QUANTIZER[0], 1);                       //  store  1e0  in G_DECQUAD_QUANTIZER[0]

  assert( decQuadDigits(     &G_DECQUAD_QUANTIZER[0]) == 1 );
  assert( decQuadGetExponent(&G_DECQUAD_QUANTIZER[0]) == 0 );

  for ( i=1; i<=DECQUAD_Pmax; i++ ) {                                 // in G_DECQUAD_QUANTIZER[1..DECQUAD_Pmax]
      decQuadCopy(&G_DECQUAD_QUANTIZER[i], &G_DECQUAD_QUANTIZER[0]);

      decQuadSetExponent(&G_DECQUAD_QUANTIZER[i], &set, -i);          // store 1e-1 .. 1e-DECQUAD_Pmax
  }

  assert( decQuadDigits(     &G_DECQUAD_QUANTIZER[DECQUAD_Pmax]) == 1 );
  assert( decQuadGetExponent(&G_DECQUAD_QUANTIZER[DECQUAD_Pmax]) == -DECQUAD_Pmax );  // -34


  //----- fill G_DECQUAD_INTEGRAL_PART_QUANTIZER[] -----

  decQuadFromInt32(&G_DECQUAD_INTEGRAL_PART_QUANTIZER[0], 1);              //  store  1e0  in G_DECQUAD_INTEGRAL_PART_QUANTIZER[0]

  assert( decQuadDigits(     &G_DECQUAD_INTEGRAL_PART_QUANTIZER[0]) == 1 );
  assert( decQuadGetExponent(&G_DECQUAD_INTEGRAL_PART_QUANTIZER[0]) == 0 );

  for ( i=1; i<=DECQUAD_Pmax+1; i++ ) {                                    // in G_DECQUAD_INTEGRAL_PART_QUANTIZER[1..DECQUAD_Pmax+1]
      decQuadCopy(&G_DECQUAD_INTEGRAL_PART_QUANTIZER[i], &G_DECQUAD_INTEGRAL_PART_QUANTIZER[0]);

      decQuadSetExponent(&G_DECQUAD_INTEGRAL_PART_QUANTIZER[i], &set, i);  // store 1e1 .. 1e(DECQUAD_Pmax+1)
  }

  assert( decQuadDigits(     &G_DECQUAD_INTEGRAL_PART_QUANTIZER[DECQUAD_Pmax])   == 1 );
  assert( decQuadGetExponent(&G_DECQUAD_INTEGRAL_PART_QUANTIZER[DECQUAD_Pmax])   == DECQUAD_Pmax   );  // 34

  assert( decQuadDigits(     &G_DECQUAD_INTEGRAL_PART_QUANTIZER[DECQUAD_Pmax+1]) == 1 );
  assert( decQuadGetExponent(&G_DECQUAD_INTEGRAL_PART_QUANTIZER[DECQUAD_Pmax+1]) == DECQUAD_Pmax+1 );  // 35


  //----- check for errors or any warning -----

  if ( set.status ) {
      s = decContextStatusToString(&set);
      fprintf(stderr, "INITIALIZATION mydecquad.c:mdq_init() FAILED: decNumber quantizer initialization failed. %s\n", s);
      exit(1);
  }

}


/************************************************************************/
/*                        arithmetic operations                         */
/************************************************************************/


/* returns 0E0.
*/
decQuad mdq_zero() {
  decQuad  val;

  decQuadZero(&val);

  return val;
}

/* returns NaN.
*/
decQuad mdq_nan() {
  decContext set;
  decQuad    val;

  decContextDefault(&set, DEC_INIT_DECQUAD);

  decQuadFromString(&val, "Nan", &set);

  //assert(set.status & DEC_Errors == 0); // a status bit is set, because of Nan

  return val;
}


/* unary minus.
*/
Quad mdq_minus(Quad a) {
  decContext  set;
  Quad        res;

  decContextDefault(&set, DEC_INIT_DECQUAD);
  set.status = a.status;

  decQuadMinus(&res.val, &a.val, &set);
  res.status = decContextGetStatus(&set);

  return res;
}


/* addition.
*/
Quad mdq_add(Quad a, Quad b) {
  decContext  set;
  Quad        res;

  decContextDefault(&set, DEC_INIT_DECQUAD);
  set.status = a.status | b.status;

  decQuadAdd(&res.val, &a.val, &b.val, &set);
  res.status = decContextGetStatus(&set);

  return res;
}


/* subtraction.
*/
Quad mdq_subtract(Quad a, Quad b) {
  decContext  set;
  Quad        res;

  decContextDefault(&set, DEC_INIT_DECQUAD);
  set.status = a.status | b.status;

  decQuadSubtract(&res.val, &a.val, &b.val, &set);
  res.status = decContextGetStatus(&set);

  return res;
}


/* multiplication.
*/
Quad mdq_multiply(Quad a, Quad b) {
  decContext  set;
  Quad        res;

  decContextDefault(&set, DEC_INIT_DECQUAD);
  set.status = a.status | b.status;

  decQuadMultiply(&res.val, &a.val, &b.val, &set);
  res.status = decContextGetStatus(&set);

  return res;
}


/* division.
*/
Quad mdq_divide(Quad a, Quad b) {
  decContext  set;
  Quad        res;

  decContextDefault(&set, DEC_INIT_DECQUAD);
  set.status = a.status | b.status;

  decQuadDivide(&res.val, &a.val, &b.val, &set);
  res.status = decContextGetStatus(&set);

  return res;
}


/* integer division.
*/
Quad mdq_divide_integer(Quad a, Quad b) {
  decContext  set;
  Quad        res;

  decContextDefault(&set, DEC_INIT_DECQUAD);
  set.status = a.status | b.status;

  decQuadDivideInteger(&res.val, &a.val, &b.val, &set);
  res.status = decContextGetStatus(&set);

  return res;
}


/* modulo.
*/
Quad mdq_remainder(Quad a, Quad b) {
  decContext  set;
  Quad        res;

  decContextDefault(&set, DEC_INIT_DECQUAD);
  set.status = a.status | b.status;

  decQuadRemainder(&res.val, &a.val, &b.val, &set);
  res.status = decContextGetStatus(&set);

  return res;
}


/* max.
*/
Quad mdq_max(Quad a, Quad b) {
  decContext  set;
  Quad        res;

  decContextDefault(&set, DEC_INIT_DECQUAD);
  set.status = a.status | b.status;

  decQuadMax(&res.val, &a.val, &b.val, &set);
  res.status = decContextGetStatus(&set);

  return res;
}


/* min.
*/
Quad mdq_min(Quad a, Quad b) {
  decContext  set;
  Quad        res;

  decContextDefault(&set, DEC_INIT_DECQUAD);
  set.status = a.status | b.status;

  decQuadMin(&res.val, &a.val, &b.val, &set);
  res.status = decContextGetStatus(&set);

  return res;
}


/* to integral value.
*/
Quad mdq_to_integral(Quad a, int round) {
  decContext  set;
  Quad        res;

  decContextDefault(&set, DEC_INIT_DECQUAD);
  set.status = a.status;

  decQuadToIntegralValue(&res.val, &a.val, &set, round); // The DEC_Inexact flag is not set by this function, even if rounding ocurred.
  res.status = decContextGetStatus(&set);

  return res;
}


/* quantize.
*/
Quad mdq_quantize(Quad a, Quad b, int round) {
  decContext  set;
  Quad        res;

  decContextDefault(&set, DEC_INIT_DECQUAD);
  decContextSetRounding(&set, round);      // change rounding mode
  set.status = a.status | b.status;

  decQuadQuantize(&res.val, &a.val, &b.val, &set);
  res.status = decContextGetStatus(&set);

  return res;
}


/* absolute value.
*/
Quad mdq_abs(Quad a) {
  decContext  set;
  Quad        res;

  decContextDefault(&set, DEC_INIT_DECQUAD);
  set.status = a.status;

  decQuadAbs(&res.val, &a.val, &set);
  res.status = decContextGetStatus(&set);

  return res;
}


/************************************************************************/
/*                           is_finite, etc                             */
/************************************************************************/


/* check if a is Finite number.
*/
uint32_t mdq_is_finite(decQuad a) {

  return decQuadIsFinite(&a);
}


/* check if a is integer number.
*/
uint32_t mdq_is_integer(decQuad a) {

  return decQuadIsInteger(&a);
}


/* check if a is Infinite.
*/
uint32_t mdq_is_infinite(decQuad a) {

  return decQuadIsInfinite(&a);
}


/* check if a is Nan.
*/
uint32_t mdq_is_nan(decQuad a) {

  return decQuadIsNaN(&a);
}


/* check if a is > 0 and not Nan.
*/
uint32_t mdq_is_positive(decQuad a) {

  return decQuadIsPositive(&a);
}


/* check if a is == 0.
*/
uint32_t mdq_is_zero(decQuad a) {

  return decQuadIsZero(&a);
}


/* check if a is < 0 and not Nan.
*/
uint32_t mdq_is_negative(decQuad a) {

  return decQuadIsNegative(&a);
}


/* get exponent.
*/
int32_t mdq_get_exponent(decQuad a) {

  return decQuadGetExponent(&a);
}


/************************************************************************/
/*                               comparison                             */
/************************************************************************/


/* compare.
*/
uint32_t mdq_compare(Quad a, Quad b) {
  decContext      set;
  decQuad         cmp_val;
  uint32_t        res;

  decContextDefault(&set, DEC_INIT_DECQUAD);
  set.status = a.status | b.status;


  decQuadCompare(&cmp_val, &a.val, &b.val, &set); // result may be –1, 0, 1, or NaN. NaN is returned only if a or b is a NaN.

  if ( decQuadIsNaN(&cmp_val) ) {
      return CMP_NAN;
  }

  if ( decQuadIsZero(&cmp_val) ) {
      return CMP_EQUAL;
  }

  if ( decQuadIsPositive(&cmp_val) ) {
      return CMP_GREATER;
  }

  assert( decQuadIsNegative(&cmp_val) );

  return CMP_LESS;
}


/************************************************************************/
/*                    conversion from string or numbers                 */
/************************************************************************/


/* conversion from string.
*/
Quad mdq_from_string(char *s) {
  decContext  set;
  Quad        res;

  decContextDefault(&set, DEC_INIT_DECQUAD);

  decQuadFromString(&res.val, s, &set);
  res.status = decContextGetStatus(&set);

  return res;
}


/* conversion from int32.
*/
Quad mdq_from_int32(int32_t value) {
  Quad        res;

  decQuadFromInt32(&res.val, value);
  res.status = 0; // never fails

  return res;
}


/* conversion from int64.
*/
Quad mdq_from_int64(int64_t value) {
  char         buff[30]; // more than enough to store a int64     max val: 9,223,372,036,854,775,807
  decContext  set;
  Quad        res;

  decContextDefault(&set, DEC_INIT_DECQUAD);

  sprintf(buff, "%lld", (long long int)value); // write value into buffer

  decQuadFromString(&res.val, buff, &set);     // raises an error if string is invalid
  res.status = decContextGetStatus(&set);

  return res;
}


/************************************************************************/
/*                        conversion to string                          */
/************************************************************************/


/* write decQuad into byte array.

   A terminating 0 is written in the array.
   Never fails.

   The function decQuadToString() uses exponential notation too often in my opinion. E.g. 0.0000001 returns "1E-7".

   Unlike the original function decQuadToString, this function discards '-' sign for negative zero: -0.00 is displayed as "0.00".
*/
Ret_str mdq_QuadToString(decQuad a) {

  decQuad  a_aux; // copy of a, but with negative sign discarded if a is negative zero
  Ret_str  res = {.length = 0};

  if ( decQuadIsZero(&a) ) {
      decQuadCopyAbs(&a_aux, &a); // discard '-' sign if any
      decQuadToString(&a_aux, res.s);
  } else {
      decQuadToString(&a, res.s);
  }

  res.length = strlen(res.s);

  return res;
}


/* write decQuad into BCD_array.

   The returned fields are:
      BCD:       byte array. The coefficient is written one digit per byte.
      exp:       if a is not Inf or Nan, will contain the exponent.
      sign:      if negative and not zero, sign bit is set.
                 THE SIGN IS VALID ALSO IF THE FUNCTION RETURNS MDQ_INFINITE, so that we can know if it is +Inf or -Inf.
*/
Ret_BCD mdq_to_BCD(decQuad a) {

  int32_t     exp;
  uint32_t    sign;
  Ret_BCD     res = {.inf_nan = 0, .exp = 0, .sign = 0};

  // convert to BCD

  decQuadToBCD(&a, &exp, res.BCD);  // this function returns a sign bit, but we don't use it because we don't want -0

  sign = decQuadIsNegative(&a);     // 0 is never negative


  // check that result is not Inf nor Nan

  if ( ! decQuadIsFinite(&a) ) {
      if ( decQuadIsInfinite(&a) ) {
          res.inf_nan = MDQ_INFINITE;
      } else {
          res.inf_nan = MDQ_NAN;
      }
      return res;
  }

  res.exp      = exp;
  res.sign     = sign;

  return res;
}


/************************************************************************/
/*                         conversion to numbers                        */
/************************************************************************/


/* convert decQuad to int32_t
*/
Ret_int32_t mdq_to_int32(Quad a, int round) {
  decContext      set;
  Ret_int32_t     res;

  decContextDefault(&set, DEC_INIT_DECQUAD);

  res.val = decQuadToInt32(&a.val, &set, round);
  res.status = decContextGetStatus(&set);

  return res;
}


/* convert decQuad to int64_t
*/
Ret_int64_t mdq_to_int64(Quad a, int round) {
  decContext   set;
  decQuad      a_integral;
  decQuad      a_integral_quantized;
  char         a_str[DECQUAD_String];
  char        *tailptr;
  int64_t      r_val;
  Ret_int64_t  res;

  decContextDefault(&set, DEC_INIT_DECQUAD);

  decQuadToIntegralValue(&a_integral, &a.val, &set, round); // rounds the number to an integral. Only numbers with exponent<0 are rounded and shifted so that exponent becomes 0.

  decQuadQuantize(&a_integral_quantized, &a_integral, &static_one, &set); // for numbers with exponent>0. E.g. change 1e3 to 1000

  if (set.status & DEC_Errors) {
    res.status = decContextGetStatus(&set);
    res.val = 0;
    return res;
  }

  if (! decQuadIsFinite(&a_integral_quantized)) {
    decContextSetStatus(&set, DEC_Invalid_operation);
    res.status = decContextGetStatus(&set);
    res.val = 0;
    return res;
  }

  assert(decQuadGetExponent(&a_integral_quantized) == 0); // in the absence of decQuadQuantize error, the exponent of the result is always equal to that of the model 'static_one'

  decQuadToString(&a_integral_quantized, a_str);  // never raises error. Exponential notation never occurs for integral, which allows strtoll() to parse the number.
  //printf("xxxxxxxxxxxxxx  %s\n", a_str);

  errno = 0;
  r_val = strtoll(a_str, &tailptr, 10);  // changes errno if error

  if ( errno ) { // in particular, if a_str is an integer that overflows int64
    decContextSetStatus(&set, DEC_Invalid_operation);
    res.status = decContextGetStatus(&set);
    res.val = 0;
    return res;
  }

  if ( *tailptr != 0 ) { // may happen for e.g.  123e10, because it parses up to 'e'
    decContextSetStatus(&set, DEC_Invalid_operation);
    res.status = decContextGetStatus(&set);
    res.val = 0;
    return res;
  }

  res.status = decContextGetStatus(&set);
  res.val = r_val;
  return res;
}


/************************************************************************/
/*                       rounding and truncating                        */
/************************************************************************/


Quad mdq_roundM(Quad a, int32_t n, int round) {
  decContext        set;
  decQuad           r;
  decQuad          *operation_quantizer;
  Quad              res;


  decContextDefault(&set, DEC_INIT_DECQUAD);
  set.status = a.status;


  // if n is out-of-range, return Invalid_operation

  if ( n > 34 || n < -35 ) {
      decContextSetStatus(&set, DEC_Invalid_operation); // add flag to status

      res.val = mdq_nan();
      res.status = decContextGetStatus(&set);
      return res;
  }


  // operation

  decContextSetRounding(&set, round);                           // change rounding mode

  if ( n >= 0 ) {   // round or truncate fractional part
      operation_quantizer = &G_DECQUAD_QUANTIZER[n];                   // n is [0..34]

      decQuadQuantize(&res.val, &a.val, operation_quantizer, &set);        // rounding, e.g. quaantize(1234.5678, 2)  --> 1234.57

  } else {          // n < 0, round or truncate integral part
      operation_quantizer = &G_DECQUAD_INTEGRAL_PART_QUANTIZER[-n];    // -n is [0..35]

      decQuadQuantize(&r, &a.val, operation_quantizer, &set);              // rounding, e.g. quaantize(1234.5678, -2) --> 12E2
      decQuadQuantize(&res.val, &r, &G_DECQUAD_QUANTIZER[0], &set);    // right-shift the number, adding missing 0s on the left. E.g. 12E2 --> 1200E0
  }


  res.status = decContextGetStatus(&set);

  return res;
}


