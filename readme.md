# Litxap

This is a proof of concept for a stress-finder for inflected Na'vi words.
It cannot be used out of the box, as it needs a dictionary implementation,
like [litxap-fwew](https://github.com/gissleh/litxap-fwew) which is based on the Fwew Go library.

It handles ambiguities by returning every valid match, and leaves it up to
the calling code to figure out how to deal with them.

```golang
dictionary := litxap.MultiDictionary{
	litxapfwew.Global(),
	litxapfwew.MultiWordPartDictionary(),
	&litxap.NumberDictionary{},
}

input := "Kaltxì, ma kifkey!"
res, err := litxap.RunLine(line, dictionary)
if err != nil {
	log.Fatalln("error running litxap:", err)
}
```
