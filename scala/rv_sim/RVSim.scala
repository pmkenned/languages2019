import scala.io.Source
import scala.collection.mutable.ArrayBuffer
import scala.collection.mutable.Map

object Registers extends Enumeration {
    type Register = Value
    val R_X0 = Value
    val R_RA = Value
    val R_SP = Value
    val R_GP = Value
    val R_TP = Value
    val R_T0 = Value
    val R_T1 = Value
    val R_T2 = Value
    val R_S0 = Value
    //val R_FP = Value(8)
    val R_S1 = Value
    val R_A0 = Value
    val R_A1 = Value
    val R_A2 = Value
    val R_A3 = Value
    val R_A4 = Value
    val R_A5 = Value
    val R_A6 = Value
    val R_A7 = Value
    val R_S2 = Value
    val R_S3 = Value
    val R_S4 = Value
    val R_S5 = Value
    val R_S6 = Value
    val R_S7 = Value
    val R_S8 = Value
    val R_S9 = Value
    val R_S10 = Value
    val R_S11 = Value
    val R_T3 = Value
    val R_T4 = Value
    val R_T5 = Value
    val R_T6 = Value
}

object RVSim {

    private var regs = new Array[Int](32)
    private var pc: Int = 0
    private var memory = new scala.collection.immutable.HashMap[Int,Long]()

    def hexToLong(s: String): Long = {
        val s_no_0x = {
            if (s.length > 1 && s.substring(0,2) == "0x") {
                s.substring(2)
            } else {
                s
            }
        }
        s_no_0x.toList.map("0123456789abcdef".indexOf(_).toLong).reduceLeft(_ * 16 + _)
    }

    def M_r(addr: Int, size: Int): Int = {
        var data: Int = 0
        val _addr = addr % 1024
        //if (size == 1) {
        //    uint32_t x = memory[addr/4];
        //    if      (addr % 4 == 0) { data = (x & 0xff); }
        //    else if (addr % 4 == 1) { data = (x & 0xff00) >> 8; }
        //    else if (addr % 4 == 2) { data = (x & 0xff0000) >> 16; }
        //    else if (addr % 4 == 3) { data = (x & 0xff000000) >> 24; }
        //} else if (size == 2) {
        //    uint32_t x = memory[addr/4];
        //    if      (addr % 4 == 0) { data = (x & 0xffff); }
        //    else if (addr % 4 == 2) { data = (x & 0xffff0000) >> 16; }
        //    else {
        //        fprintf(stderr, "error: unaligned read access, addr = %08x, size = %d\n", addr, size);
        //        assert(0);
        //    }
        //} else if (size == 4) {
        //    if (addr % 4 != 0) {
        //        fprintf(stderr, "error: unaligned read access, addr = %08x, size = %d\n", addr, size);
        //    }
        //    assert(addr % 4 == 0);
        //    data = memory[addr/4];
        //}
        return data
    }

    def M_w(addr: Int, data: Int, size: Int): Unit = {
    }

    // TODO
    def sext(x: Int, w: Int): Int = {
        return 0
    }

    // TODO
    // def get_imm(i_p: instr_t *): Int = {
    // }

    // TODO
    //def snprint_instr(char * str, size_t size, instr_t * i_p): Int = {
    //}

    // TODO
    //def decode(op: uint32_t): instr = {
    //}

    def reset(): Unit = {

        for (i <- 0 until regs.length) {
            regs(i) = 0
        }
        pc = 0

        // TODO: remove these; put them in the test itself
        regs(Registers.R_A1.id) = 5    // n 
        regs(Registers.R_A0.id) = 0x54 // &a[0]
        regs(Registers.R_RA.id) = 0x4c // return address
    }

    def step(): Boolean = {
        return false
    }

    def print_state() = {
        println("")
        //regs.foreach(x => println(s"x.: $x"))
        for (i <- 0 until regs.length) {
            println(s"x$i: ${regs(i)}")
        }
        println(s"pc: $pc")
        println("memory:")
        for (i <- 0 to 64) {
            println(s"M[${i*16}]: ${M_r(i*16, 4)} ${M_r(i*16+4, 4)} ${M_r(i*16+8, 4)} ${M_r(i*16+12, 4)}")
        }
        println("")
    }

    def main(args: Array[String]) = {
        //val filename = args(0)
        val filename = "insert_sort.hex"
        val ops = ArrayBuffer[Long]()
        //val memory = Map[Int,Long]()
        for (line <- Source.fromFile(filename).getLines) {
            ops.append(hexToLong(line))
        }

        var i = 0
        for (op <- ops) {
            memory(i) = ops(i)
            i += 1
        }

        reset()

        print_state()
        var step_cnt = 0
        var done = false
        val MAX_STEPS = 100
        while (!done && (step_cnt < MAX_STEPS)) {
            done = step()
            step_cnt += 1
        }
        if (step_cnt >= MAX_STEPS) {
            System.err.println("exceeded MAX_STEPS")
        }
        print_state()
    }
}
