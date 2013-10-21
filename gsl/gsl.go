/* Package gsl wraps the GNU Scientific Library
 */
package gsl

import (
	"fmt"
	"math"
)

type Errno int

func (e Errno) Error() string {
	s := errText[e]
	if s == "" {
		return fmt.Sprintf("errno %d", int(e))
	}
	return s
}

// GSL Error codes. This is based on GSL v1.15, and could
// go out of date at some point.
var (
	GSL_CONTINUE = Errno(-2) /* iteration has not converged */
	GSL_FAILURE  = Errno(-1)
	GSL_SUCCESS  = Errno(0)
	GSL_EDOM     = Errno(1)  /* input domain error, e.g sqrt(-1) */
	GSL_ERANGE   = Errno(2)  /* output range error, e.g. exp(1e100) */
	GSL_EFAULT   = Errno(3)  /* invalid pointer */
	GSL_EINVAL   = Errno(4)  /* invalid argument supplied by user */
	GSL_EFAILED  = Errno(5)  /* generic failure */
	GSL_EFACTOR  = Errno(6)  /* factorization failed */
	GSL_ESANITY  = Errno(7)  /* sanity check failed - shouldn't happen */
	GSL_ENOMEM   = Errno(8)  /* malloc failed */
	GSL_EBADFUNC = Errno(9)  /* problem with user-supplied function */
	GSL_ERUNAWAY = Errno(10) /* iterative process is out of control */
	GSL_EMAXITER = Errno(11) /* exceeded max number of iterations */
	GSL_EZERODIV = Errno(12) /* tried to divide by zero */
	GSL_EBADTOL  = Errno(13) /* user specified an invalid tolerance */
	GSL_ETOL     = Errno(14) /* failed to reach the specified tolerance */
	GSL_EUNDRFLW = Errno(15) /* underflow */
	GSL_EOVRFLW  = Errno(16) /* overflow  */
	GSL_ELOSS    = Errno(17) /* loss of accuracy */
	GSL_EROUND   = Errno(18) /* failed because of roundoff error */
	GSL_EBADLEN  = Errno(19) /* matrix, vector lengths are not conformant */
	GSL_ENOTSQR  = Errno(20) /* matrix not square */
	GSL_ESING    = Errno(21) /* apparent singularity detected */
	GSL_EDIVERGE = Errno(22) /* integral or series is divergent */
	GSL_EUNSUP   = Errno(23) /* requested feature is not supported by the hardware */
	GSL_EUNIMPL  = Errno(24) /* requested feature not (yet) implemented */
	GSL_ECACHE   = Errno(25) /* cache limit exceeded */
	GSL_ETABLE   = Errno(26) /* table limit exceeded */
	GSL_ENOPROG  = Errno(27) /* iteration is not making progress towards solution */
	GSL_ENOPROGJ = Errno(28) /* jacobian evaluations are not improving the solution */
	GSL_ETOLF    = Errno(29) /* cannot reach the specified tolerance in F */
	GSL_ETOLX    = Errno(30) /* cannot reach the specified tolerance in X */
	GSL_ETOLG    = Errno(31) /* cannot reach the specified tolerance in gradient */
	GSL_EOF      = Errno(32) /* end of file */
)

var errText = map[Errno]string{
	Errno(-2): "iteration has not converged",
	Errno(-1): "GSL Failure",
	Errno(0):  "Success",
	Errno(1):  "input domain error, e.g sqrt(-1) ",
	Errno(2):  "output range error, e.g. exp(1e100) ",
	Errno(3):  "invalid pointer ",
	Errno(4):  "invalid argument supplied by user ",
	Errno(5):  "generic failure ",
	Errno(6):  "factorization failed ",
	Errno(7):  "sanity check failed - shouldn't happen ",
	Errno(8):  "malloc failed ",
	Errno(9):  "problem with user-supplied function ",
	Errno(10): "iterative process is out of control ",
	Errno(11): "exceeded max number of iterations ",
	Errno(12): "tried to divide by zero ",
	Errno(13): "user specified an invalid tolerance ",
	Errno(14): "failed to reach the specified tolerance ",
	Errno(15): "underflow ",
	Errno(16): "overflow  ",
	Errno(17): "loss of accuracy ",
	Errno(18): "failed because of roundoff error ",
	Errno(19): "matrix, vector lengths are not conformant ",
	Errno(20): "matrix not square ",
	Errno(21): "apparent singularity detected ",
	Errno(22): "integral or series is divergent ",
	Errno(23): "requested feature is not supported by the hardware ",
	Errno(24): "requested feature not (yet) implemented ",
	Errno(25): "cache limit exceeded ",
	Errno(26): "table limit exceeded ",
	Errno(27): "iteration is not making progress towards solution ",
	Errno(28): "jacobian evaluations are not improving the solution ",
	Errno(29): "cannot reach the specified tolerance in F ",
	Errno(30): "cannot reach the specified tolerance in X ",
	Errno(31): "cannot reach the specified tolerance in gradient ",
	Errno(32): "end of file",
}

// Function type
type F func(float64) float64

// Wrappers for Evalers
type GSLFuncWrapper struct {
	Gofunc F
}

// Eps packages the tolerances
type Eps struct {
	Abs, Rel float64
}

// Result packages the results
type Result struct {
	Res, Err float64
}

// Interval packages an Interval
type Interval struct {
	Lo, Hi float64
}

// Define infinities
var (
	Inf  = math.Inf(1)
	NInf = math.Inf(-1)
)
