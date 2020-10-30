// to test how multiple go files in a single folder

package main
import ff "common/package1/Foo"

func main() {
	Bar()
	ff.FooOne()
	ff.FooTwo()
}