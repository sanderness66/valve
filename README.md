# valve(1) - decode thermionic valve designations

Kozmix Go, 01-SEP-2022

```
valve [name ...]
```


<a name="description"></a>

# Description

**valve**
will run through valve names given either on the command line or as
stdin and attempt to decode any information that can be gleaned from
them.


<a name="example"></a>

# Example

.EX
$ valve ecc83
ECC83
heater:         6.3V or 6.3/12.6V
base:           B9A noval
function:       1. Small-signal triode
function:       2. Small-signal triode
aliases:        B339 7025 12AX7 6057 M8137

.EE


<a name="bugs"></a>

# Bugs

**valve**
works best with European Philips/Mullard designations (like ECC83) or
American R.E.T.M.A. designations for which
**valve**
knows a Philips/Mullard equivalent (like 12AX7). It can tell very
little from R.E.T.M.A. names alone, and nothing at all from valve
designations that don't follow either scheme.

**valve**
will sometimes mis-identify valves that don't use either of the two
naming schemes mentioned (for example the KT66, which looks like a
Philips/Mullard name to
**valve**
but isn't).

**valve**
hasn't been hipped yet to some very obscure valve sockets
(Philips/Mullard names that end in more than two digits) or to
inverted names (like E83CC).


<a name="author"></a>

# Author

svm
