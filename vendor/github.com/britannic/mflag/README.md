# mflag

`import "github.com/britannic/mflag"`

- [Overview](#pkg-overview)
- [Index](#pkg-index)
- [Examples](#pkg-examples)

## <a name="pkg-overview">Overview</a>

## <a name="pkg-index">Index</a>

- [Variables](#pkg-variables)
- [func Arg(i int) string](#Arg)
- [func Args() \[\]string](#Args)
- [func Bool(name string, value bool, usage string) \*bool](#Bool)
- [func BoolVar(p \*bool, name string, value bool, usage string)](#BoolVar)
- [func Duration(name string, value time.Duration, usage string) \*time.Duration](#Duration)
- [func DurationVar(p \*time.Duration, name string, value time.Duration, usage string)](#DurationVar)
- [func Float64(name string, value float64, usage string) \*float64](#Float64)
- [func Float64Var(p \*float64, name string, value float64, usage string)](#Float64Var)
- [func Int(name string, value int, usage string) \*int](#Int)
- [func Int64(name string, value int64, usage string) \*int64](#Int64)
- [func Int64Var(p \*int64, name string, value int64, usage string)](#Int64Var)
- [func IntVar(p \*int, name string, value int, usage string)](#IntVar)
- [func IsZeroValue(flag \*Flag, value string) bool](#IsZeroValue)
- [func NArg() int](#NArg)
- [func NFlag() int](#NFlag)
- [func Parse()](#Parse)
- [func Parsed() bool](#Parsed)
- [func PrintDefaults()](#PrintDefaults)
- [func Set(name, value string) error](#Set)
- [func String(name string, value string, usage string) \*string](#String)
- [func StringVar(p \*string, name string, value string, usage string)](#StringVar)
- [func Uint(name string, value uint, usage string) \*uint](#Uint)
- [func Uint64(name string, value uint64, usage string) \*uint64](#Uint64)
- [func Uint64Var(p \*uint64, name string, value uint64, usage string)](#Uint64Var)
- [func UintVar(p \*uint, name string, value uint, usage string)](#UintVar)
- [func UnquoteUsage(flag \*Flag) (name string, usage string)](#UnquoteUsage)
- [func Var(value Value, name string, usage string)](#Var)
- [func Visit(fn func(\*Flag))](#Visit)
- [func VisitAll(fn func(\*Flag))](#VisitAll)
- [type ErrorHandling](#ErrorHandling)
- [type Flag](#Flag)
  - [func Lookup(name string) \*Flag](#Lookup)
- [type FlagSet](#FlagSet)
  - [[func NewFlagSet(name string, errorHandling ErrorHandling) \*FlagSet](#NewFlagSet)
  - [[func (f \*FlagSet) Arg(i int) string](#FlagSet.Arg)
  - [[func (f \*FlagSet) Args() \[\]string](#FlagSet.Args)
  - [[func (f _FlagSet) Bool(name string, value bool, usage string) _bool](#FlagSet.Bool)
  - [[func (f _FlagSet) BoolVar(p _bool, name string, value bool, usage string)](#FlagSet.BoolVar)
  - [[func (f _FlagSet) Duration(name string, value time.Duration, usage string) _time.Duration](#FlagSet.Duration)
  - [[func (f _FlagSet) DurationVar(p _time.Duration, name string, value time.Duration, usage string)](#FlagSet.DurationVar)
  - [[func (f _FlagSet) Float64(name string, value float64, usage string) _float64](#FlagSet.Float64)
  - [[func (f _FlagSet) Float64Var(p _float64, name string, value float64, usage string)](#FlagSet.Float64Var)
  - [[func (f \*FlagSet) Init(name string, errorHandling ErrorHandling)](#FlagSet.Init)
  - [[func (f _FlagSet) Int(name string, value int, usage string) _int](#FlagSet.Int)
  - [[func (f _FlagSet) Int64(name string, value int64, usage string) _int64](#FlagSet.Int64)
  - [[func (f _FlagSet) Int64Var(p _int64, name string, value int64, usage string)](#FlagSet.Int64Var)
  - [[func (f _FlagSet) IntVar(p _int, name string, value int, usage string)](#FlagSet.IntVar)
  - [[func (f _FlagSet) Lookup(name string) _Flag](#FlagSet.Lookup)
  - [[func (f \*FlagSet) NArg() int](#FlagSet.NArg)
  - [[func (f \*FlagSet) NFlag() int](#FlagSet.NFlag)
  - [[func (f \*FlagSet) Out() io.Writer](#FlagSet.Out)
  - [[func (f \*FlagSet) Parse(arguments \[\]string) error](#FlagSet.Parse)
  - [[func (f \*FlagSet) Parsed() bool](#FlagSet.Parsed)
  - [[func (f \*FlagSet) PrintDefaults()](#FlagSet.PrintDefaults)
  - [[func (f \*FlagSet) Set(name, value string) error](#FlagSet.Set)
  - [[func (f \*FlagSet) SetOutput(output io.Writer)](#FlagSet.SetOutput)
  - [[func (f _FlagSet) String(name string, value string, usage string) _string](#FlagSet.String)
  - [[func (f _FlagSet) StringVar(p _string, name string, value string, usage string)](#FlagSet.StringVar)
  - [[func (f _FlagSet) Uint(name string, value uint, usage string) _uint](#FlagSet.Uint)
  - [[func (f _FlagSet) Uint64(name string, value uint64, usage string) _uint64](#FlagSet.Uint64)
  - [[func (f _FlagSet) Uint64Var(p _uint64, name string, value uint64, usage string)](#FlagSet.Uint64Var)
  - [[func (f _FlagSet) UintVar(p _uint, name string, value uint, usage string)](#FlagSet.UintVar)
  - [[func (f \*FlagSet) Var(value Value, name string, usage string)](#FlagSet.Var)
  - [[func (f _FlagSet) Visit(fn func(_Flag))](#FlagSet.Visit)
  - [[func (f _FlagSet) VisitAll(fn func(_Flag))](#FlagSet.VisitAll)
- [type Getter](#Getter)
- [type StringValue](#StringValue)
  - [[func NewStringValue(val string, p _string) _StringValue](#NewStringValue)
  - [[func (s \*StringValue) Get() interface{}](#StringValue.Get)
  - [[func (s \*StringValue) Set(val string) error](#StringValue.Set)
  - [[func (s \*StringValue) String() string](#StringValue.String)
- [type Value](#Value)

### <a name="pkg-examples">Examples</a>

- [Package](#example_)

### <a name="pkg-files">Package files</a>

[mflag.go](/src/github.com/britannic/mflag/mflag.go)

## <a name="pkg-variables">Variables</a>

```go
var CommandLine = NewFlagSet(os.Args[0], ExitOnError)
```

CommandLine is the default set of command-line flags, parsed from os.Args.
The top-level functions such as BoolVar, Arg, and so on are wrappers for the
methods of CommandLine.

```go
var ErrHelp = errors.New("flag: help requested")
```

ErrHelp is the error returned if the -help or -h flag is invoked
but no such flag is defined.

```go
var Usage = func() {
    fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
    PrintDefaults()
}
```

Usage prints to standard error a usage message documenting all defined command-line flags.
It is called when an error occurs while parsing flags.
The function is a variable that may be changed to point to a custom function.
By default it prints a simple header and calls PrintDefaults; for details about the
format of the output and how to control it, see the documentation for PrintDefaults.

## <a name="Arg">func</a> [Arg](/src/target/mflag.go?s=16234:16256#L540)

```go
func Arg(i int) string
```

Arg returns the i'th command-line argument. Arg(0) is the first remaining argument
after flags have been processed. Arg returns an empty string if the
requested element does not exist.

## <a name="Args">func</a> [Args](/src/target/mflag.go?s=16694:16714#L554)

```go
func Args() []string
```

Args returns the non-flag command-line arguments.

## <a name="Bool">func</a> [Bool](/src/target/mflag.go?s=17802:17856#L578)

```go
func Bool(name string, value bool, usage string) *bool
```

Bool defines a bool flag with specified name, default value, and usage string.
The return value is the address of a bool variable that stores the value of the flag.

## <a name="BoolVar">func</a> [BoolVar](/src/target/mflag.go?s=17207:17267#L564)

```go
func BoolVar(p *bool, name string, value bool, usage string)
```

BoolVar defines a bool flag with specified name, default value, and usage string.
The argument p points to a bool variable in which to store the value of the flag.

## <a name="Duration">func</a> [Duration](/src/target/mflag.go?s=26588:26664#L764)

```go
func Duration(name string, value time.Duration, usage string) *time.Duration
```

Duration defines a time.Duration flag with specified name, default value, and usage string.
The return value is the address of a time.Duration variable that stores the value of the flag.
The flag accepts a value acceptable to time.ParseDuration.

## <a name="DurationVar">func</a> [DurationVar](/src/target/mflag.go?s=25764:25846#L748)

```go
func DurationVar(p *time.Duration, name string, value time.Duration, usage string)
```

DurationVar defines a time.Duration flag with specified name, default value, and usage string.
The argument p points to a time.Duration variable in which to store the value of the flag.
The flag accepts a value acceptable to time.ParseDuration.

## <a name="Float64">func</a> [Float64](/src/target/mflag.go?s=24990:25053#L734)

```go
func Float64(name string, value float64, usage string) *float64
```

Float64 defines a float64 flag with specified name, default value, and usage string.
The return value is the address of a float64 variable that stores the value of the flag.

## <a name="Float64Var">func</a> [Float64Var](/src/target/mflag.go?s=24350:24419#L720)

```go
func Float64Var(p *float64, name string, value float64, usage string)
```

Float64Var defines a float64 flag with specified name, default value, and usage string.
The argument p points to a float64 variable in which to store the value of the flag.

## <a name="Int">func</a> [Int](/src/target/mflag.go?s=18948:18999#L604)

```go
func Int(name string, value int, usage string) *int
```

Int defines an int flag with specified name, default value, and usage string.
The return value is the address of an int variable that stores the value of the flag.

## <a name="Int64">func</a> [Int64](/src/target/mflag.go?s=20140:20197#L630)

```go
func Int64(name string, value int64, usage string) *int64
```

Int64 defines an int64 flag with specified name, default value, and usage string.
The return value is the address of an int64 variable that stores the value of the flag.

## <a name="Int64Var">func</a> [Int64Var](/src/target/mflag.go?s=19526:19589#L616)

```go
func Int64Var(p *int64, name string, value int64, usage string)
```

Int64Var defines an int64 flag with specified name, default value, and usage string.
The argument p points to an int64 variable in which to store the value of the flag.

## <a name="IntVar">func</a> [IntVar](/src/target/mflag.go?s=18364:18421#L590)

```go
func IntVar(p *int, name string, value int, usage string)
```

IntVar defines an int flag with specified name, default value, and usage string.
The argument p points to an int variable in which to store the value of the flag.

## <a name="IsZeroValue">func</a> [IsZeroValue](/src/target/mflag.go?s=10823:10870#L377)

```go
func IsZeroValue(flag *Flag, value string) bool
```

IsZeroValue guesses whether the string represents the zero
value for a flag. It is not accurate but in practice works OK.

## <a name="NArg">func</a> [NArg](/src/target/mflag.go?s=16498:16513#L548)

```go
func NArg() int
```

NArg is the number of arguments remaining after flags have been processed.

## <a name="NFlag">func</a> [NFlag](/src/target/mflag.go?s=15700:15716#L525)

```go
func NFlag() int
```

NFlag returns the number of command-line flags that have been set.

## <a name="Parse">func</a> [Parse](/src/target/mflag.go?s=31586:31598#L932)

```go
func Parse()
```

Parse parses the command-line flags from os.Args[1:].  Must be called
after all flags are defined and before flags are accessed by the program.

## <a name="Parsed">func</a> [Parsed](/src/target/mflag.go?s=31758:31776#L938)

```go
func Parsed() bool
```

Parsed reports whether the command-line flags have been parsed.

## <a name="PrintDefaults">func</a> [PrintDefaults](/src/target/mflag.go?s=14562:14582#L493)

```go
func PrintDefaults()
```

PrintDefaults prints, to standard error unless configured otherwise,
a usage message showing the default settings of all defined
command-line flags.
For an integer valued flag x, the default output has the form

    -x int
        usage-message-for-x (default 7)

The usage message will appear on a separate line for anything but
a bool flag with a one-byte name. For bool flags, the type is
omitted and if the flag name is one byte the usage message appears
on the same line. The parenthetical default is omitted if the
default is the zero value for the type. The listed type, here int,
can be changed by placing a back-quoted name in the flag's usage
string; the first such item in the message is taken to be a parameter
name to show in the message and the back quotes are stripped from
the message when displayed. For instance, given

    mflag.String("I", "", "search `directory` for include files")

the output will be

    -I directory
        search directory for include files.

## <a name="Set">func</a> [Set](/src/target/mflag.go?s=10618:10652#L371)

```go
func Set(name, value string) error
```

Set sets the value of the named command-line flag.

## <a name="String">func</a> [String](/src/target/mflag.go?s=23744:23804#L708)

```go
func String(name string, value string, usage string) *string
```

String defines a string flag with specified name, default value, and usage string.
The return value is the address of a string variable that stores the value of the flag.

## <a name="StringVar">func</a> [StringVar](/src/target/mflag.go?s=23119:23185#L694)

```go
func StringVar(p *string, name string, value string, usage string)
```

StringVar defines a string flag with specified name, default value, and usage string.
The argument p points to a string variable in which to store the value of the flag.

## <a name="Uint">func</a> [Uint](/src/target/mflag.go?s=21310:21364#L656)

```go
func Uint(name string, value uint, usage string) *uint
```

Uint defines a uint flag with specified name, default value, and usage string.
The return value is the address of a uint  variable that stores the value of the flag.

## <a name="Uint64">func</a> [Uint64](/src/target/mflag.go?s=22523:22583#L682)

```go
func Uint64(name string, value uint64, usage string) *uint64
```

Uint64 defines a uint64 flag with specified name, default value, and usage string.
The return value is the address of a uint64 variable that stores the value of the flag.

## <a name="Uint64Var">func</a> [Uint64Var](/src/target/mflag.go?s=21898:21964#L668)

```go
func Uint64Var(p *uint64, name string, value uint64, usage string)
```

Uint64Var defines a uint64 flag with specified name, default value, and usage string.
The argument p points to a uint64 variable in which to store the value of the flag.

## <a name="UintVar">func</a> [UintVar](/src/target/mflag.go?s=20713:20773#L642)

```go
func UintVar(p *uint, name string, value uint, usage string)
```

UintVar defines a uint flag with specified name, default value, and usage string.
The argument p points to a uint  variable in which to store the value of the flag.

## <a name="UnquoteUsage">func</a> [UnquoteUsage](/src/target/mflag.go?s=11733:11790#L408)

```go
func UnquoteUsage(flag *Flag) (name string, usage string)
```

UnquoteUsage extracts a back-quoted name from the usage
string for a flag and returns it and the un-quoted usage.
Given "a `name` to show" it returns ("name", "a name to show").
If there are no back quotes, the name is an educated guess of the
type of the flag's value, or the empty string if the flag is boolean.

## <a name="Var">func</a> [Var](/src/target/mflag.go?s=28203:28251#L800)

```go
func Var(value Value, name string, usage string)
```

Var defines a flag with the specified name and usage string. The type and
value of the flag are represented by the first argument, of type Value, which
typically holds a user-defined implementation of Value. For instance, the
caller could create a flag that turns a comma-separated string into a slice
of strings by giving the slice the methods of Value; in particular, Set would
decompose the comma-separated string into the slice.

## <a name="Visit">func</a> [Visit](/src/target/mflag.go?s=9840:9866#L338)

```go
func Visit(fn func(*Flag))
```

Visit visits the command-line flags in lexicographical order, calling fn
for each. It visits only those flags that have been set.

## <a name="VisitAll">func</a> [VisitAll](/src/target/mflag.go?s=9417:9446#L324)

```go
func VisitAll(fn func(*Flag))
```

VisitAll visits the command-line flags in lexicographical order, calling
fn for each. It visits all flags, even those not set.

## <a name="ErrorHandling">type</a> [ErrorHandling](/src/target/mflag.go?s=7210:7232#L250)

```go
type ErrorHandling int
```

ErrorHandling defines how FlagSet.Parse behaves if the parse fails.

```go
const (
    ContinueOnError ErrorHandling = iota // Return a descriptive error.
    ExitOnError                          // Call os.Exit(2).
    PanicOnError                         // Call panic with a descriptive error.
)
```

These constants cause FlagSet.Parse to behave as described if the parse fails.

## <a name="Flag">type</a> [Flag](/src/target/mflag.go?s=8172:8376#L277)

```go
type Flag struct {
    Name     string // name as it appears on command line
    Usage    string // help message
    Value    Value  // value as set
    DefValue string // default value (as text); for usage message
}
```

A Flag represents the state of a flag.

### <a name="Lookup">func</a> [Lookup](/src/target/mflag.go?s=10155:10185#L349)

```go
func Lookup(name string) *Flag
```

Lookup returns the Flag structure of the named command-line flag,
returning nil if none exists.

## <a name="FlagSet">type</a> [FlagSet](/src/target/mflag.go?s=7663:8128#L261)

```go
type FlagSet struct {
    // Usage is the function called when an error occurs while parsing flags.
    // The field is a function (not a method) that may be changed to point to
    // a custom error handler.
    Usage func()

    Output io.Writer // nil means stderr; use Out() accessor
    // contains filtered or unexported fields
}
```

A FlagSet represents a set of defined flags. The zero value of a FlagSet
has no name and has ContinueOnError error handling.

### <a name="NewFlagSet">func</a> [NewFlagSet](/src/target/mflag.go?s=32484:32550#L961)

```go
func NewFlagSet(name string, errorHandling ErrorHandling) *FlagSet
```

NewFlagSet returns a new, empty flag set with the specified name and
error handling property.

### <a name="FlagSet.Arg">func</a> (\*FlagSet) [Arg](/src/target/mflag.go?s=15934:15969#L530)

```go
func (f *FlagSet) Arg(i int) string
```

Arg returns the i'th argument. Arg(0) is the first remaining argument
after flags have been processed. Arg returns an empty string if the
requested element does not exist.

### <a name="FlagSet.Args">func</a> (\*FlagSet) [Args](/src/target/mflag.go?s=16588:16621#L551)

```go
func (f *FlagSet) Args() []string
```

Args returns the non-flag arguments.

### <a name="FlagSet.Bool">func</a> (\*FlagSet) [Bool](/src/target/mflag.go?s=17498:17565#L570)

```go
func (f *FlagSet) Bool(name string, value bool, usage string) *bool
```

Bool defines a bool flag with specified name, default value, and usage string.
The return value is the address of a bool variable that stores the value of the flag.

### <a name="FlagSet.BoolVar">func</a> (\*FlagSet) [BoolVar](/src/target/mflag.go?s=16914:16987#L558)

```go
func (f *FlagSet) BoolVar(p *bool, name string, value bool, usage string)
```

BoolVar defines a bool flag with specified name, default value, and usage string.
The argument p points to a bool variable in which to store the value of the flag.

### <a name="FlagSet.Duration">func</a> (\*FlagSet) [Duration](/src/target/mflag.go?s=26165:26254#L755)

```go
func (f *FlagSet) Duration(name string, value time.Duration, usage string) *time.Duration
```

Duration defines a time.Duration flag with specified name, default value, and usage string.
The return value is the address of a time.Duration variable that stores the value of the flag.
The flag accepts a value acceptable to time.ParseDuration.

### <a name="FlagSet.DurationVar">func</a> (\*FlagSet) [DurationVar](/src/target/mflag.go?s=25361:25456#L741)

```go
func (f *FlagSet) DurationVar(p *time.Duration, name string, value time.Duration, usage string)
```

DurationVar defines a time.Duration flag with specified name, default value, and usage string.
The argument p points to a time.Duration variable in which to store the value of the flag.
The flag accepts a value acceptable to time.ParseDuration.

### <a name="FlagSet.Float64">func</a> (\*FlagSet) [Float64](/src/target/mflag.go?s=24662:24738#L726)

```go
func (f *FlagSet) Float64(name string, value float64, usage string) *float64
```

Float64 defines a float64 flag with specified name, default value, and usage string.
The return value is the address of a float64 variable that stores the value of the flag.

### <a name="FlagSet.Float64Var">func</a> (\*FlagSet) [Float64Var](/src/target/mflag.go?s=24036:24118#L714)

```go
func (f *FlagSet) Float64Var(p *float64, name string, value float64, usage string)
```

Float64Var defines a float64 flag with specified name, default value, and usage string.
The argument p points to a float64 variable in which to store the value of the flag.

### <a name="FlagSet.Init">func</a> (\*FlagSet) [Init](/src/target/mflag.go?s=32833:32897#L973)

```go
func (f *FlagSet) Init(name string, errorHandling ErrorHandling)
```

Init sets the name and error handling property for a flag set.
By default, the zero FlagSet uses an empty name and the
ContinueOnError error handling policy.

### <a name="FlagSet.Int">func</a> (\*FlagSet) [Int](/src/target/mflag.go?s=18650:18714#L596)

```go
func (f *FlagSet) Int(name string, value int, usage string) *int
```

Int defines an int flag with specified name, default value, and usage string.
The return value is the address of an int variable that stores the value of the flag.

### <a name="FlagSet.Int64">func</a> (\*FlagSet) [Int64](/src/target/mflag.go?s=19826:19896#L622)

```go
func (f *FlagSet) Int64(name string, value int64, usage string) *int64
```

Int64 defines an int64 flag with specified name, default value, and usage string.
The return value is the address of an int64 variable that stores the value of the flag.

### <a name="FlagSet.Int64Var">func</a> (\*FlagSet) [Int64Var](/src/target/mflag.go?s=19224:19300#L610)

```go
func (f *FlagSet) Int64Var(p *int64, name string, value int64, usage string)
```

Int64Var defines an int64 flag with specified name, default value, and usage string.
The argument p points to an int64 variable in which to store the value of the flag.

### <a name="FlagSet.IntVar">func</a> (\*FlagSet) [IntVar](/src/target/mflag.go?s=18076:18146#L584)

```go
func (f *FlagSet) IntVar(p *int, name string, value int, usage string)
```

IntVar defines an int flag with specified name, default value, and usage string.
The argument p points to an int variable in which to store the value of the flag.

### <a name="FlagSet.Lookup">func</a> (\*FlagSet) [Lookup](/src/target/mflag.go?s=9981:10024#L343)

```go
func (f *FlagSet) Lookup(name string) *Flag
```

Lookup returns the Flag structure of the named flag, returning nil if none exists.

### <a name="FlagSet.NArg">func</a> (\*FlagSet) [NArg](/src/target/mflag.go?s=16367:16395#L545)

```go
func (f *FlagSet) NArg() int
```

NArg is the number of arguments remaining after flags have been processed.

### <a name="FlagSet.NFlag">func</a> (\*FlagSet) [NFlag](/src/target/mflag.go?s=15574:15603#L522)

```go
func (f *FlagSet) NFlag() int
```

NFlag returns the number of flags that have been set.

### <a name="FlagSet.Out">func</a> (\*FlagSet) [Out](/src/target/mflag.go?s=8784:8817#L301)

```go
func (f *FlagSet) Out() io.Writer
```

Out returns a non nil io.Writer for log output

### <a name="FlagSet.Parse">func</a> (\*FlagSet) [Parse](/src/target/mflag.go?s=30992:31041#L902)

```go
func (f *FlagSet) Parse(arguments []string) error
```

Parse parses flag definitions from the argument list, which should not
include the command name. Must be called after all flags in the FlagSet
are defined and before flags are accessed by the program.
The return value will be ErrHelp if -help or -h were set but not defined.

### <a name="FlagSet.Parsed">func</a> (\*FlagSet) [Parsed](/src/target/mflag.go?s=31382:31413#L926)

```go
func (f *FlagSet) Parsed() bool
```

Parsed reports whether f.Parse has been called.

### <a name="FlagSet.PrintDefaults">func</a> (\*FlagSet) [PrintDefaults](/src/target/mflag.go?s=12708:12741#L445)

```go
func (f *FlagSet) PrintDefaults()
```

PrintDefaults prints to standard error the default values of all
defined command-line flags in the set. See the documentation for
the global function PrintDefaults for more information.

### <a name="FlagSet.Set">func</a> (\*FlagSet) [Set](/src/target/mflag.go?s=10265:10312#L354)

```go
func (f *FlagSet) Set(name, value string) error
```

Set sets the value of the named flag.

### <a name="FlagSet.SetOutput">func</a> (\*FlagSet) [SetOutput](/src/target/mflag.go?s=8988:9033#L310)

```go
func (f *FlagSet) SetOutput(output io.Writer)
```

SetOutput sets the destination for usage and error messages.
If output is nil, os.Stderr is used.

### <a name="FlagSet.String">func</a> (\*FlagSet) [String](/src/target/mflag.go?s=23424:23497#L700)

```go
func (f *FlagSet) String(name string, value string, usage string) *string
```

String defines a string flag with specified name, default value, and usage string.
The return value is the address of a string variable that stores the value of the flag.

### <a name="FlagSet.StringVar">func</a> (\*FlagSet) [StringVar](/src/target/mflag.go?s=22812:22891#L688)

```go
func (f *FlagSet) StringVar(p *string, name string, value string, usage string)
```

StringVar defines a string flag with specified name, default value, and usage string.
The argument p points to a string variable in which to store the value of the flag.

### <a name="FlagSet.Uint">func</a> (\*FlagSet) [Uint](/src/target/mflag.go?s=21005:21072#L648)

```go
func (f *FlagSet) Uint(name string, value uint, usage string) *uint
```

Uint defines a uint flag with specified name, default value, and usage string.
The return value is the address of a uint  variable that stores the value of the flag.

### <a name="FlagSet.Uint64">func</a> (\*FlagSet) [Uint64](/src/target/mflag.go?s=22203:22276#L674)

```go
func (f *FlagSet) Uint64(name string, value uint64, usage string) *uint64
```

Uint64 defines a uint64 flag with specified name, default value, and usage string.
The return value is the address of a uint64 variable that stores the value of the flag.

### <a name="FlagSet.Uint64Var">func</a> (\*FlagSet) [Uint64Var](/src/target/mflag.go?s=21591:21670#L662)

```go
func (f *FlagSet) Uint64Var(p *uint64, name string, value uint64, usage string)
```

Uint64Var defines a uint64 flag with specified name, default value, and usage string.
The argument p points to a uint64 variable in which to store the value of the flag.

### <a name="FlagSet.UintVar">func</a> (\*FlagSet) [UintVar](/src/target/mflag.go?s=20419:20492#L636)

```go
func (f *FlagSet) UintVar(p *uint, name string, value uint, usage string)
```

UintVar defines a uint flag with specified name, default value, and usage string.
The argument p points to a uint variable in which to store the value of the flag.

### <a name="FlagSet.Var">func</a> (\*FlagSet) [Var](/src/target/mflag.go?s=27170:27231#L774)

```go
func (f *FlagSet) Var(value Value, name string, usage string)
```

Var defines a flag with the specified name and usage string. The type and
value of the flag are represented by the first argument, of type Value, which
typically holds a user-defined implementation of Value. For instance, the
caller could create a flag that turns a comma-separated string into a slice
of strings by giving the slice the methods of Value; in particular, Set would
decompose the comma-separated string into the slice.

### <a name="FlagSet.Visit">func</a> (\*FlagSet) [Visit](/src/target/mflag.go?s=9601:9640#L330)

```go
func (f *FlagSet) Visit(fn func(*Flag))
```

Visit visits the flags in lexicographical order, calling fn for each.
It visits only those flags that have been set.

### <a name="FlagSet.VisitAll">func</a> (\*FlagSet) [VisitAll](/src/target/mflag.go?s=9178:9220#L316)

```go
func (f *FlagSet) VisitAll(fn func(*Flag))
```

VisitAll visits the flags in lexicographical order, calling fn for each.
It visits all flags, even those not set.

## <a name="Getter">type</a> [Getter](/src/target/mflag.go?s=7086:7137#L244)

```go
type Getter interface {
    Value
    Get() interface{}
}
```

Getter is an interface that allows the contents of a Value to be retrieved.
It wraps the Value interface, rather than being part of it, because it
appeared after Go 1 and its compatibility rules. All Value types provided
by this package satisfy the Getter interface.

## <a name="StringValue">type</a> [StringValue](/src/target/mflag.go?s=4925:4948#L170)

```go
type StringValue string
```

StringValue is a string type

### <a name="NewStringValue">func</a> [NewStringValue](/src/target/mflag.go?s=5003:5058#L173)

```go
func NewStringValue(val string, p *string) *StringValue
```

NewStringValue returns a pointer to a StringValue

### <a name="StringValue.Get">func</a> (\*StringValue) [Get](/src/target/mflag.go?s=5272:5311#L185)

```go
func (s *StringValue) Get() interface{}
```

Get returns the string value of a StringValue

### <a name="StringValue.Set">func</a> (\*StringValue) [Set](/src/target/mflag.go?s=5139:5182#L179)

```go
func (s *StringValue) Set(val string) error
```

Set sets the value of a StringValue

### <a name="StringValue.String">func</a> (\*StringValue) [String](/src/target/mflag.go?s=5335:5372#L187)

```go
func (s *StringValue) String() string
```

## <a name="Value">type</a> [Value](/src/target/mflag.go?s=6745:6805#L235)

```go
type Value interface {
    String() string
    Set(string) error
}
```

Value is the interface to the dynamic value stored in a flag.
(The default value is represented as a string.)

If a Value has an IsBoolFlag() bool method returning true,
the command-line parser makes -name equivalent to -name=true
rather than using the next command-line argument.

Set is called once, in command line order, for each flag present.
The flag package may call the String method with a zero-valued receiver,
such as a nil pointer.

* * *

## Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
