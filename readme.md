# __gotk3-builder__

## _Streamlining the use of builder files_

This package is designed to simplify the loading of builder files by replacing this code:

```go 
//get the builder object 	
obj, err := builder.GetObject("window")
if err != nil {
    log.Fatalf("%v not found in builder file: %v", item, err.Error())
}
// convert it to the correct widget
if win, ok := obj.(*gtk.Window); !ok {
	log.Fatalf("%v is not a gtk.Window", item)
} 
// the variable "win" now contains the "window" widget pointer
```

with one line:

```go 
win := ui.GetWindow("window")
```

## __Limitations:__

Only one builder file is supported. The builder object is stored in a local package variable so that it will not need to
be passed by the calling function. However, this means that all widgets are available to any part of the program that
imports the "gotk3-builder" package.

There is no recovery from errors, all errors fatal out. This should not cause an issue as any error in the builder file
that prevents a widget from loading is usually a fatal error. This is done to allow the following to be done in 2 lines
rather than 16 lines:

```go
text := ui.GetEntry("name_field").GetText()
ui.GetEntry("address").SetText("123 Main Street")
```

## __Usage__:

### Load the builder file:

```go
import ui github.com/graynerd/gotk3-builder
...
ui.LoadBuilder("~/app/builder.ui")
```

The "builder.ui" file is now loaded into the package and all "Get..." calls will load widgets from this file.

### Load the various widgets:

Import the package into every file that will access the builder file. A alias can be used here to shorten the name.

```go
import ui github.com/graynerd/gotk3-builder
```

The name that is passed as the parameter is the ID from the required widget.

```go 
win := ui.GetWindow("main_window")

field1 := GetEntry("name")

combo := GetComboBox("combo")

button := GetButton("quit_button")
```

Any function from the widget can be appended to the "Get..." call:

```go
field1 := GetEntry("name")
field1.SetText("John Doe")
```

becomes

```go
GetEntry("name").SetText("John Doe")
```

The example applications shows a basic application load that demostrates the functionality.

### Future

Not all widgets are implemented. More will be added in the future. However it is quite easy to add your own. Just append
following to the ui.go file and change the <WIDGET> to the name of the GTK+ Widget (i.e. GetEntryText, GetComboBoxText,
...)

```go
    // Get<WIDGET> returns a pointer to the named gtk.xxxxx
func Get<WIDGET>(item string) *gtk.<WIDGET> {
if obj, err := builder.GetObject(item); err == nil {
if e, ok := obj.(*gtk.<WIDGET>; ok {
return e
}
log.Fatalf("%v is not a *gtk.<WIDGET>", item)
} else {
log.Fatalf("%v not found in builder file: %v", item, err.Error())
}
return nil
}
```

 



