trait Greeter {
    def greet(name: String): Unit =
        println("Hello, " + name + "!")
}

class DefaultGreeter() extends Greeter

class CustomizableGreeter(prefix: String, suffix: String) extends Greeter {
    override def greet(name: String): Unit =
        println(prefix + name + suffix)
}

// instances of case classes are immutable by default
// they are compared by value
case class Point(x: Int, y: Int)

// objects are singletons of their own classes
object IdFactory {
    private var counter = 0
    def create(): Int = {
        counter += 1
        counter
    }
}

object Hello {
    def main(args: Array[String]) = {
        val x : Int = 2
        // type can be inferred
        var y = 2
        // vars can be reassigned
        y = 3
        println("hello, " + x + y)
        println(x + y)

        // can combine expressions with {}
        println({
            val x = 7
            var y = 7
            x + y
        }) // 14

        // anonymous function
        val add = (x: Int, y: Int) => x + y
        println(add(1,2))

        // method
        def addThenMul(x: Int, y: Int)(z: Int): Int  = (x + y)*z
        println(addThenMul(2,1)(3))

        def getSquareString(input: Double): String = {
            val square = input * input
            square.toString
        }

        println(getSquareString(5.0))

        val greeter = new DefaultGreeter()
        greeter.greet("Paul")

        val customGreeter = new CustomizableGreeter("How are you, ", "?")
        customGreeter.greet("Paul")

        val point1 = Point(1, 2)
        val point2 = Point(1, 2)
        val point3 = Point(2, 2)
        if (point1 == point2) {
            println(s"$point1 and $point2 are the same")
        }
        if (point1 != point3) {
            println(s"$point1 and $point3 are different")
        }

        val id1 = IdFactory.create()

        // types which are subtypes of AnyVal
        val d: Double   = 1.0
        val f: Float    = 1.0f
        val l: Long     = 100
        val i: Int      = 1
        val s: Short    = 1
        val b: Byte     = 1
        val c: Char     = 'a'
        val u: Unit     = ()
        val o: Boolean  = true
        val a: Any      = 3
        val v: AnyVal   = 3

        // AnyRef

        val list: List[Any] = List(
            "a string",
            732,
            'c',
            true,
            () => "anonymous function returning a string"
        )

        println(list)

        // Type casting:
        // Byte->Short->Int->Long->Float->Double
        //              ^
        //              Char

        // Nothing: a subtype of all types; there is no value with type Nothing
        // Null: a subtype of all reference types; has a single value: null

        class User

        val user = new User

    }
}
