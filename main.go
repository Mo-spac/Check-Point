package main

import (
	problems_2 "check_point/problems_2"
	problems_3 "check_point/problems_3"
	"check_point/problems_4"

	// problems_4 "check_point/problems_4"
	"fmt"
)

func main() {

	fmt.Println(problems_2.Checknumber("Hello"))
	fmt.Println(problems_2.Checknumber("Hello world"))
	fmt.Println(problems_2.Checknumber("Hello1"))
	fmt.Println(problems_2.Checknumber("Hello15"))
	fmt.Println(problems_2.Checknumber("Hello-8"))
	fmt.Println(problems_2.Checknumber("Hello9"))
	fmt.Println(problems_2.Checknumber("12345"))
	fmt.Println(problems_2.Checknumber(""))
	fmt.Println(problems_2.Checknumber(" "))

	println("***************")

	fmt.Println(problems_2.CountChar("Hello World", 'l'))
	fmt.Println(problems_2.CountChar("5  balloons", 5))
	fmt.Println(problems_2.CountChar("   ", ' '))
	fmt.Println(problems_2.CountChar("The 7 deadly sins", '7'))

	println("***************")
	fmt.Println(problems_2.Countalpha("Hello world"))
	fmt.Println(problems_2.Countalpha("H e l l o"))
	fmt.Println(problems_2.Countalpha("H1e2l3l4o"))

	println("***************")

	fmt.Print(problems_2.Printif("abcdefz"))
	fmt.Print(problems_2.Printif("abc"))
	fmt.Print(problems_2.Printif(""))
	fmt.Print(problems_2.Printif("14"))

	println("***************")

	fmt.Print(problems_2.PrintIfNot("abcdefz"))
	fmt.Print(problems_2.PrintIfNot("abc"))
	fmt.Print(problems_2.PrintIfNot(""))
	fmt.Print(problems_2.PrintIfNot("14"))

	println("***************")

	fmt.Println(problems_2.RectPerimeter(10, 2))
	fmt.Println(problems_2.RectPerimeter(434343, 898989))
	fmt.Println(problems_2.RectPerimeter(10, -2))

	println("***************")

	fmt.Println(problems_2.RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(problems_2.RetainFirstHalf("A"))
	fmt.Println(problems_2.RetainFirstHalf(""))
	fmt.Println(problems_2.RetainFirstHalf("Hello World"))

	println("***************")

	fmt.Println(problems_3.CamelToSnakeCase("HelloWorld"))       // الناتج المتوقع: Hello_World
	fmt.Println(problems_3.CamelToSnakeCase("helloWorld"))       // الناتج المتوقع: hello_World
	fmt.Println(problems_3.CamelToSnakeCase("helloWorlD"))       // الناتج المتوقع: helloWorlD (ينتهي بحرف كبير)
	fmt.Println(problems_3.CamelToSnakeCase("camelCase"))        // الناتج المتوقع: camel_Case
	fmt.Println(problems_3.CamelToSnakeCase("CamelCasE"))        // الناتج المتوقع: CamelCasE (ينتهي بحرف كبير)
	fmt.Println(problems_3.CamelToSnakeCase("CamelCAse"))        // الناتج المتوقع: CamelCAse (حرفان كبيران متتاليان مع وجود حروف صغيرة)
	fmt.Println(problems_3.CamelToSnakeCase("CAMELtoSnackCASE")) // الناتج المتوقع: CAMELtoSnackCASE (غير صالح)
	fmt.Println(problems_3.CamelToSnakeCase("camelToSnakeCase")) // الناتج المتوقع: camel_To_Snake_Case
	fmt.Println(problems_3.CamelToSnakeCase("hey2"))             // الناتج المتوقع: hey2 (يحتوي على رقم)

	println("***************")

	fmt.Println(problems_3.CamelToSnakeCase1("HelloWorld"))       // الناتج المتوقع: Hello_World
	fmt.Println(problems_3.CamelToSnakeCase1("helloWorld"))       // الناتج المتوقع: hello_World
	fmt.Println(problems_3.CamelToSnakeCase1("helloWorlD"))       // الناتج المتوقع: helloWorlD (ينتهي بحرف كبير)
	fmt.Println(problems_3.CamelToSnakeCase1("camelCase"))        // الناتج المتوقع: camel_Case
	fmt.Println(problems_3.CamelToSnakeCase1("CamelCasE"))        // الناتج المتوقع: CamelCasE (ينتهي بحرف كبير)
	fmt.Println(problems_3.CamelToSnakeCase1("CamelCAse"))        // الناتج المتوقع: CamelCAse (حرفان كبيران متتاليان مع وجود حروف صغيرة)
	fmt.Println(problems_3.CamelToSnakeCase1("CAMELtoSnackCASE")) // الناتج المتوقع: CAMELtoSnackCASE (غير صالح)
	fmt.Println(problems_3.CamelToSnakeCase1("camelToSnakeCase")) // الناتج المتوقع: camel_To_Snake_Case
	fmt.Println(problems_3.CamelToSnakeCase1("hey2"))             // الناتج المتوقع: hey2 (يحتوي على رقم)

	println("***************")

	fmt.Println(problems_3.CamelToSnakeCase2("HelloWorld"))       // الناتج المتوقع: Hello_World
	fmt.Println(problems_3.CamelToSnakeCase2("helloWorld"))       // الناتج المتوقع: hello_World
	fmt.Println(problems_3.CamelToSnakeCase2("helloWorlD"))       // الناتج المتوقع: helloWorlD (ينتهي بحرف كبير)
	fmt.Println(problems_3.CamelToSnakeCase2("camelCase"))        // الناتج المتوقع: camel_Case
	fmt.Println(problems_3.CamelToSnakeCase2("CamelCasE"))        // الناتج المتوقع: CamelCasE (ينتهي بحرف كبير)
	fmt.Println(problems_3.CamelToSnakeCase2("CamelCAse"))        // الناتج المتوقع: CamelCAse (حرفان كبيران متتاليان مع وجود حروف صغيرة)
	fmt.Println(problems_3.CamelToSnakeCase2("CAMELtoSnackCASE")) // الناتج المتوقع: CAMELtoSnackCASE (غير صالح)
	fmt.Println(problems_3.CamelToSnakeCase2("camelToSnakeCase")) // الناتج المتوقع: camel_To_Snake_Case
	fmt.Println(problems_3.CamelToSnakeCase2("hey2"))             // الناتج المتوقع: hey2 (يحتوي على رقم)

	println("***************")

	fmt.Println(problems_3.CamelToSnakeCase3("HelloWorld"))       // الناتج المتوقع: Hello_World
	fmt.Println(problems_3.CamelToSnakeCase3("helloWorld"))       // الناتج المتوقع: hello_World
	fmt.Println(problems_3.CamelToSnakeCase3("helloWorlD"))       // الناتج المتوقع: helloWorlD (ينتهي بحرف كبير)
	fmt.Println(problems_3.CamelToSnakeCase3("camelCase"))        // الناتج المتوقع: camel_Case
	fmt.Println(problems_3.CamelToSnakeCase3("CamelCasE"))        // الناتج المتوقع: CamelCasE (ينتهي بحرف كبير)
	fmt.Println(problems_3.CamelToSnakeCase3("CamelCAse"))        // الناتج المتوقع: CamelCAse (حرفان كبيران متتاليان مع وجود حروف صغيرة)
	fmt.Println(problems_3.CamelToSnakeCase3("CAMELtoSnackCASE")) // الناتج المتوقع: CAMELtoSnackCASE (غير صالح)
	fmt.Println(problems_3.CamelToSnakeCase3("camelToSnakeCase")) // الناتج المتوقع: camel_To_Snake_Case
	fmt.Println(problems_3.CamelToSnakeCase3("hey2"))             // الناتج المتوقع: hey2 (يحتوي على رقم)

	println("***************")

	fmt.Println(problems_3.CamelToSnakeCase4("HelloWorld"))       // الناتج المتوقع: Hello_World
	fmt.Println(problems_3.CamelToSnakeCase4("helloWorld"))       // الناتج المتوقع: hello_World
	fmt.Println(problems_3.CamelToSnakeCase4("helloWorlD"))       // الناتج المتوقع: helloWorlD (ينتهي بحرف كبير)
	fmt.Println(problems_3.CamelToSnakeCase4("camelCase"))        // الناتج المتوقع: camel_Case
	fmt.Println(problems_3.CamelToSnakeCase4("CamelCasE"))        // الناتج المتوقع: CamelCasE (ينتهي بحرف كبير)
	fmt.Println(problems_3.CamelToSnakeCase4("CamelCAse"))        // الناتج المتوقع: CamelCAse (حرفان كبيران متتاليان مع وجود حروف صغيرة)
	fmt.Println(problems_3.CamelToSnakeCase4("CAMELtoSnackCASE")) // الناتج المتوقع: CAMELtoSnackCASE (غير صالح)
	fmt.Println(problems_3.CamelToSnakeCase4("camelToSnakeCase")) // الناتج المتوقع: camel_To_Snake_Case
	fmt.Println(problems_3.CamelToSnakeCase4("hey2"))             // الناتج المتوقع: hey2 (يحتوي على رقم)

	println("***************************")

	fmt.Println(problems_3.DigitLen(100, 10))
	fmt.Println(problems_3.DigitLen(100, 2))
	fmt.Println(problems_3.DigitLen(-100, 16))
	fmt.Println(problems_3.DigitLen(100, -1))

	println("***************")

	fmt.Println(problems_3.FishAndChips(4))
	fmt.Println(problems_3.FishAndChips(9))
	fmt.Println(problems_3.FishAndChips(6))
	fmt.Println(problems_3.FishAndChips(-6))

	println("***************")

	print(problems_3.FirstWord("hello there"))
	print(problems_3.FirstWord(" "))
	print(problems_3.FirstWord("hello .......... bye"))
	print(problems_3.FirstWord("hello"))
	print(problems_3.FirstWord("   hello there"))

	println("***************")

	print(problems_3.FirstWord1("hello there"))
	print(problems_3.FirstWord1(" "))
	print(problems_3.FirstWord1("hello .......... bye"))
	print(problems_3.FirstWord1("hello"))
	print(problems_3.FirstWord1("   hello there"))

	println("***************")

	print(problems_3.LastWord("hello there"))
	print(problems_3.LastWord(" "))
	print(problems_3.LastWord("hello .......... bye"))
	print(problems_3.LastWord("hello"))
	print(problems_3.LastWord("   hello there"))

	println("***************")

	print(problems_3.LastWord("this        ...       is sparta, then again, maybe    not"))
	print(problems_3.LastWord(" lorem,ipsum "))
	print(problems_3.LastWord(" "))

	println("***************")

	print(problems_3.LastWord1("this        ...       is sparta, then again, maybe    not"))
	print(problems_3.LastWord1(" lorem,ipsum "))
	print(problems_3.LastWord1(" "))

	println("***************")

	fmt.Println(problems_3.Gcd(42, 10))
	fmt.Println(problems_3.Gcd(42, 12))
	fmt.Println(problems_3.Gcd(14, 77))
	fmt.Println(problems_3.Gcd(17, 3))

	println("***************")

	fmt.Println(problems_3.Gcd1(42, 10))
	fmt.Println(problems_3.Gcd1(42, 12))
	fmt.Println(problems_3.Gcd1(14, 77))
	fmt.Println(problems_3.Gcd1(17, 3))

	println("***************")

	fmt.Println(problems_3.Gcd2(42, 10))
	fmt.Println(problems_3.Gcd2(42, 12))
	fmt.Println(problems_3.Gcd2(14, 77))
	fmt.Println(problems_3.Gcd2(17, 3))

	println("***************")

	fmt.Println(problems_4.IsCapitalized("Hello! How are you?"))
	fmt.Println(problems_4.IsCapitalized("Hello How Are You"))
	fmt.Println(problems_4.IsCapitalized("Whats 4this 100K?"))
	fmt.Println(problems_4.IsCapitalized("Whatsthis4"))
	fmt.Println(problems_4.IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(problems_4.IsCapitalized(""))

	println("***************")

	fmt.Println(problems_4.IsCapitalized1("Hello! How are you?"))
	fmt.Println(problems_4.IsCapitalized1("Hello How Are You"))
	fmt.Println(problems_4.IsCapitalized1("Whats 4this 100K?"))
	fmt.Println(problems_4.IsCapitalized1("Whatsthis4"))
	fmt.Println(problems_4.IsCapitalized1("!!!!Whatsthis4"))
	fmt.Println(problems_4.IsCapitalized1(""))

	println("***************")

	fmt.Print(problems_4.FromTo(1, 10))
	fmt.Print(problems_4.FromTo(10, 1))
	fmt.Print(problems_4.FromTo(10, 10))
	fmt.Print(problems_4.FromTo(100, 10))

	println("***************")

	fmt.Println(problems_3.HashCode("A"))
	fmt.Println(problems_3.HashCode("AB"))
	fmt.Println(problems_3.HashCode("BAC"))
	fmt.Println(problems_3.HashCode("Hello World"))

	println("***************")

	fmt.Println(problems_3.HashCode1("A"))
	fmt.Println(problems_3.HashCode1("AB"))
	fmt.Println(problems_3.HashCode1("BAC"))
	fmt.Println(problems_3.HashCode1("Hello World"))

	println("***************")

	fmt.Println(problems_4.FindPrevPrime(5))  // 5
	fmt.Println(problems_4.FindPrevPrime(4))  // 3
	fmt.Println(problems_4.FindPrevPrime(31)) // 31
	fmt.Println(problems_4.FindPrevPrime(1))

	println("***************")

	fmt.Println(problems_4.FindPrevPrime1(5))  // 5
	fmt.Println(problems_4.FindPrevPrime1(4))  // 3
	fmt.Println(problems_4.FindPrevPrime1(31)) // 31
	fmt.Println(problems_4.FindPrevPrime1(1))

	println("***************")

	fmt.Println(problems_4.FindPrevPrime2(5))  // 5
	fmt.Println(problems_4.FindPrevPrime2(4))  // 3
	fmt.Println(problems_4.FindPrevPrime2(31)) // 31
	fmt.Println(problems_4.FindPrevPrime2(1))  // 0

	println("***************")

	fmt.Println(problems_4.Itoa(12345))
	fmt.Println(problems_4.Itoa(0))
	fmt.Println(problems_4.Itoa(-1234))
	fmt.Println(problems_4.Itoa(987654321))

	println("***************")

	problems_4.PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})

	println("***************")

	problems_4.Printrevcomb()

	println("***************")

	fmt.Print(problems_4.ThirdTimeIsACharm("123456789"))
	fmt.Print(problems_4.ThirdTimeIsACharm(""))
	fmt.Print(problems_4.ThirdTimeIsACharm("a b c d e f"))
	fmt.Print(problems_4.ThirdTimeIsACharm("12"))

	println("***************")

	fmt.Println(problems_4.WeAreUnique("foo", "boo"))
	fmt.Println(problems_4.WeAreUnique("", ""))
	fmt.Println(problems_4.WeAreUnique("abc", "def"))

	println("***************")

	fmt.Println(problems_4.WeAreUnique1("foo", "boo"))
	fmt.Println(problems_4.WeAreUnique1("", ""))
	fmt.Println(problems_4.WeAreUnique1("abc", "def"))

	fmt.Println("***************")

	fmt.Println(problems_4.ZipString("YouuungFellllas"))
	fmt.Println(problems_4.ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(problems_4.ZipString("Helloo Therre!"))

	fmt.Println("****************")

	fmt.Println(problems_4.ZipString1("YouuungFellllas"))
	fmt.Println(problems_4.ZipString1("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(problems_4.ZipString1("Helloo Therre!"))

}
