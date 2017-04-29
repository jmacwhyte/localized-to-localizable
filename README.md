## Localized-to-Localizable

This is a simple Go program that searches through a (swift or Obj-C) file of NSLocalizedStrings and generates a Localizable.strings file based on the keys, default values, and comments.

### But why?

XCode apparently doesn't really support exporting a `Localizable.strings` file when localized strings have a default value that doesn't match the key. This isn't a problem when you want your keys to be the same as your values, such as:

`key: "Click here to submit", value: "Click here to submit"`

However, what if you want your keys to be more code-friendly/platform agnostic while saving default values that represent the actual text you want to include? This script will allow you to easily generate a `Localizable.strings` file with strings that look like:

`key: "SubmitText", value: "Click here to submit"`

### How to use it
Simply define all of your strings with a key, default value, and (optional) comment like so:

```let submitText = NSLocalizedString("SubmitText", value:"Click here to submit", comment: "Submit button label")```

Run the script by calling `l2l` and it will look in the current directory for a `Strings.swift` file. You can also pass the filename you want it to parse as an argument after `l2l`.