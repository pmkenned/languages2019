import scala.io.Source
import scala.collection.mutable.ArrayBuffer
import scala.collection.mutable.Map

object Helper {
    def sext(x: Int, w: Int): Int = {
        if (w == 32) {
            x
        } else {
            if (((x >> (w-1)) & 0x1) == 1)
                x | ~((1 << w) - 1)
            else
                x & ((1 << w) - 1)
        }
    }

    def mask_msb_lsb(x: Int, msb: Int, lsb: Int): Int = {
        (x >> lsb) & ((1 << (msb-lsb+1)) - 1)
    }
}
import Helper._

object Mnemonic extends Enumeration {
    type Mnemonic   = Value
    val invalid     = Value
    val lui         = Value
    val auipc       = Value
    val jal         = Value
    val jalr        = Value
    val beq         = Value
    val bne         = Value
    val blt        = Value
    val bge        = Value
    val bltu       = Value
    val bgeu       = Value
    val lb         = Value
    val lh         = Value
    val lw         = Value
    val lbu        = Value
    val lhu        = Value
    val sb         = Value
    val sh         = Value
    val sw         = Value
    val addi       = Value
    val slti       = Value
    val sltiu      = Value
    val xori       = Value
    val ori        = Value
    val andi       = Value
    val slli       = Value
    val srli       = Value
    val srai       = Value
    val add        = Value
    val sub        = Value
    val sll        = Value
    val slt        = Value
    val sltu       = Value
    val xor        = Value
    val srl        = Value
    val sra        = Value
    val or         = Value
    val and        = Value
    val fence      = Value
    val fence_i    = Value
    val ecall      = Value
    val ebreak     = Value
    val csrrw      = Value
    val csrrs      = Value
    val csrrc      = Value
    val csrrwi     = Value
    val csrrsi     = Value
    val csrrci     = Value
}
import Mnemonic.Mnemonic

object Format extends Enumeration {
    type Format = Value
    val u = Value
    val i = Value
    val b = Value
    val j = Value
    val s = Value
    val r = Value
}
import Format.Format

abstract class Instr(var mnem: Mnemonic, var fmt: Format) {
    def toString(): String
    def rd(): Int
    def rs1(): Int
    def rs2(): Int
    def imm(): Int
    def shamt(): Int = { assert(false); 0 }
    //def pred(): Int = { assert(false); 0 }
    //def succ(): Int = { assert(false); 0 }
    //def csr(): Int = { assert(false); 0 }
    //def zimm(): Int = { assert(false); 0 }
}

class InstrR(mnem: Mnemonic, fmt: Format, var rd: Int, func3: Int, var rs1: Int, var rs2: Int, func7: Int) extends Instr(mnem, fmt) {
    override def toString() = {
        s"${mnem} ${rd}, ${rs1}, ${rs2}"
    }
    def imm(): Int = { assert(false); 0 }
}

class InstrI(mnem: Mnemonic, fmt: Format, var rd: Int, func3: Int, var rs1: Int, imm_11_0: Int) extends Instr(mnem, fmt) {
    override def toString() = {
        s"${mnem} ${rd}, ${rs1}, ${imm}"
    }
    def imm() = { sext(imm_11_0, 12) }
    def rs2() = { assert(false); 0 }
    override def shamt() = { mask_msb_lsb(imm_11_0, 4, 0) }
}

class InstrS(mnem: Mnemonic, fmt: Format, imm_4_0: Int, func3: Int, var rs1: Int, var rs2: Int, imm_11_5: Int) extends Instr(mnem, fmt) {
    override def toString() = {
        s"${mnem} ${rs1}, ${imm}(${rs2})"
    }
    def imm() = { sext((imm_11_5 << 5) + imm_4_0, 12) }
    def rd() = { assert(false); 0 }
}

class InstrB(mnem: Mnemonic, fmt: Format, imm_11_a7: Int, imm_4_1: Int, func3: Int, var rs1: Int, var rs2: Int, imm_12: Int, imm_10_5: Int) extends Instr(mnem, fmt) {
    override def toString() = {
        s"${mnem} ${rs1}, ${rs2}, ${imm}"
    }
    def imm() = { sext((imm_12 << 12) + (imm_11_a7 << 11) + (imm_10_5 << 5) + (imm_4_1 << 1), 13) }
    def rd() = { assert(false); 0 }
}

class InstrU(mnem: Mnemonic, fmt: Format, var rd: Int, imm_31_12: Int) extends Instr(mnem, fmt) {
    override def toString() = {
        s"${mnem} ${rd}, ${imm}"
    }
    def imm() = { imm_31_12 << 12 }
    def rs1() = { assert(false); 0 }
    def rs2() = { assert(false); 0 }
}

class InstrJ(mnem: Mnemonic, fmt: Format, var rd: Int, imm_19_12: Int, imm_11_a20: Int, imm_10_1: Int, imm_20: Int) extends Instr(mnem, fmt) {
    override def toString() = {
        s"${mnem} ${rd}, ${imm}"
    }
    def imm() = { sext((imm_20 << 20) + (imm_19_12 << 12) + (imm_11_a20 << 11) + (imm_10_1 << 1), 21) }
    def rs1() = { assert(false); 0 }
    def rs2() = { assert(false); 0 }
}

object Register extends Enumeration {
    type Register = Value
    val X0 = Value("x0")
    val RA = Value("ra")
    val SP = Value("sp")
    val GP = Value("gp")
    val TP = Value("tp")
    val T0 = Value("t0")
    val T1 = Value("t1")
    val T2 = Value("t2")
    val S0 = Value("s0")
    val S1 = Value("s1")
    val A0 = Value("a0")
    val A1 = Value("a1")
    val A2 = Value("a2")
    val A3 = Value("a3")
    val A4 = Value("a4")
    val A5 = Value("a5")
    val A6 = Value("a6")
    val A7 = Value("a7")
    val S2 = Value("s2")
    val S3 = Value("s3")
    val S4 = Value("s4")
    val S5 = Value("s5")
    val S6 = Value("s6")
    val S7 = Value("s7")
    val S8 = Value("s8")
    val S9 = Value("s9")
    val S10 = Value("s10")
    val S11 = Value("s11")
    val T3 = Value("t3")
    val T4 = Value("t4")
    val T5 = Value("t5")
    val T6 = Value("t6")
}

object RVSim {

    private var regs = new Array[Int](32)
    private var pc: Int = 0
    private val memory = Map[Int,Int]()

    def hexToInt(s: String): Int = {
        val s_no_0x = {
            if (s.length > 1 && s.substring(0,2) == "0x") {
                s.substring(2)
            } else {
                s
            }
        }
        s_no_0x.toList.map("0123456789abcdef".indexOf(_).toInt).reduceLeft(_ * 16 + _)
    }

    def M_r(addr: Int, size: Int): Int = {
        // TODO: error-handling for misaligned addresses and invalid sizes?
        assert(addr % size == 0)
        assert(size == 1 || size == 2 || size == 4)
        // TODO: random values for unitialized memory?
        val data = memory.getOrElse(addr/4, 0)
        if (size == 1) {
            data >> (8*(addr % 4)) & 0xff
        } else if (size == 2) {
            data >> (8*(addr % 4)) & 0xffff
        } else {
            data
        }
    }

    def M_w(addr: Int, data: Int, size: Int): Unit = {
        assert(addr % size == 0)
        assert(size == 1 || size == 2 || size == 4)
        val old_mask = if (size == 1) {
            if (addr % 4 == 0)      { 0xffffff00 }
            else if (addr % 4 == 1) { 0xffff00ff }
            else if (addr % 4 == 2) { 0xff00ffff }
            else                    { 0x00ffffff }
        } else if (size == 2) {
            if (addr % 4 == 0)  { 0xffff0000 }
            else                { 0x0000ffff }
        } else {
            0
        }
        val new_mask =
            if (size == 1)      { 0xff }
            else if (size == 2) { 0xffff }
            else                { 0xffffffff }
        memory(addr/4) = (M_r(addr & ~3, 4) & old_mask) | (data & new_mask) << (8*(addr % 4))
    }

    // TODO
    //def snprint_instr(char * str, size_t size, instr_t * i_p): Int = {
    //}

    // TODO
    def decode(op: Int): Instr = {

        val rd          = mask_msb_lsb(op, 11, 7)
        val rs1         = mask_msb_lsb(op, 19, 15)
        val rs2         = mask_msb_lsb(op, 24, 20)
        val func3       = mask_msb_lsb(op, 14, 12)
        val func7       = mask_msb_lsb(op, 31, 25)
        val imm_11_0    = mask_msb_lsb(op, 31, 20)
        val imm_4_0     = mask_msb_lsb(op, 11, 7)
        val imm_11_5    = mask_msb_lsb(op, 31, 25)
        val imm_11_a7   = mask_msb_lsb(op, 7, 7)
        val imm_4_1     = mask_msb_lsb(op, 11, 8)
        val imm_10_5    = mask_msb_lsb(op, 30, 25)
        val imm_12      = mask_msb_lsb(op, 31, 31)
        val imm_31_12   = mask_msb_lsb(op, 31, 12)
        val imm_19_12   = mask_msb_lsb(op, 19, 12)
        val imm_11_a20  = mask_msb_lsb(op, 20, 20)
        val imm_10_1    = mask_msb_lsb(op, 30, 21)
        val imm_20      = mask_msb_lsb(op, 31, 31)

        val mnem = if ((op & 0x0000007f) == 0x00000037) Mnemonic.lui
        else if ((op & 0x0000007f) == 0x00000017)       Mnemonic.auipc
        else if ((op & 0x0000007f) == 0x0000006f)       Mnemonic.jal
        else if ((op & 0x0000707f) == 0x00000067)       Mnemonic.jalr
        else if ((op & 0x0000707f) == 0x00000063)       Mnemonic.beq
        else if ((op & 0x0000707f) == 0x00001063)       Mnemonic.bne
        else if ((op & 0x0000707f) == 0x00004063)       Mnemonic.blt
        else if ((op & 0x0000707f) == 0x00005063)       Mnemonic.bge
        else if ((op & 0x0000707f) == 0x00006063)       Mnemonic.bltu
        else if ((op & 0x0000707f) == 0x00007063)       Mnemonic.bgeu
        else if ((op & 0x0000707f) == 0x00000003)       Mnemonic.lb
        else if ((op & 0x0000707f) == 0x00001003)       Mnemonic.lh
        else if ((op & 0x0000707f) == 0x00002003)       Mnemonic.lw
        else if ((op & 0x0000707f) == 0x00004003)       Mnemonic.lbu
        else if ((op & 0x0000707f) == 0x00005003)       Mnemonic.lhu
        else if ((op & 0x0000707f) == 0x00000023)       Mnemonic.sb
        else if ((op & 0x0000707f) == 0x00001023)       Mnemonic.sh
        else if ((op & 0x0000707f) == 0x00002023)       Mnemonic.sw
        else if ((op & 0x0000707f) == 0x00000013)       Mnemonic.addi
        else if ((op & 0x0000707f) == 0x00002013)       Mnemonic.slti
        else if ((op & 0x0000707f) == 0x00003013)       Mnemonic.sltiu
        else if ((op & 0x0000707f) == 0x00004013)       Mnemonic.xori
        else if ((op & 0x0000707f) == 0x00006013)       Mnemonic.ori
        else if ((op & 0x0000707f) == 0x00007013)       Mnemonic.andi
        else if ((op & 0xfe00707f) == 0x00001013)       Mnemonic.slli
        else if ((op & 0xfe00707f) == 0x00005013)       Mnemonic.srli
        else if ((op & 0xfe00707f) == 0x40005013)       Mnemonic.srai
        else if ((op & 0xfe00707f) == 0x00000033)       Mnemonic.add
        else if ((op & 0xfe00707f) == 0x40000033)       Mnemonic.sub
        else if ((op & 0xfe00707f) == 0x00001033)       Mnemonic.sll
        else if ((op & 0xfe00707f) == 0x00002033)       Mnemonic.slt
        else if ((op & 0xfe00707f) == 0x00003033)       Mnemonic.sltu
        else if ((op & 0xfe00707f) == 0x00004033)       Mnemonic.xor
        else if ((op & 0xfe00707f) == 0x00005033)       Mnemonic.srl
        else if ((op & 0xfe00707f) == 0x40005033)       Mnemonic.sra
        else if ((op & 0xfe00707f) == 0x00006033)       Mnemonic.or
        else if ((op & 0xfe00707f) == 0x00007033)       Mnemonic.and
        else if ((op & 0xf00fffff) == 0x0000000f)       Mnemonic.fence
        else if ((op & 0xffffffff) == 0x0000100f)       Mnemonic.fence_i
        else if ((op & 0xffffffff) == 0x00000073)       Mnemonic.ecall
        else if ((op & 0xffffffff) == 0x00100073)       Mnemonic.ebreak
        else if ((op & 0x0000707f) == 0x00001073)       Mnemonic.csrrw
        else if ((op & 0x0000707f) == 0x00002073)       Mnemonic.csrrs
        else if ((op & 0x0000707f) == 0x00003073)       Mnemonic.csrrc
        else if ((op & 0x0000707f) == 0x00005073)       Mnemonic.csrrwi
        else if ((op & 0x0000707f) == 0x00006073)       Mnemonic.csrrsi
        else if ((op & 0x0000707f) == 0x00007073)       Mnemonic.csrrci
        else Mnemonic.invalid

        val fmt = mnem match {
            case Mnemonic.lui       => Format.u
            case Mnemonic.auipc     => Format.u
            case Mnemonic.jal       => Format.j
            case Mnemonic.jalr      => Format.i
            case Mnemonic.beq       => Format.b
            case Mnemonic.bne       => Format.b
            case Mnemonic.blt       => Format.b
            case Mnemonic.bge       => Format.b
            case Mnemonic.bltu      => Format.b
            case Mnemonic.bgeu      => Format.b
            case Mnemonic.lb        => Format.i
            case Mnemonic.lh        => Format.i
            case Mnemonic.lw        => Format.i
            case Mnemonic.lbu       => Format.i
            case Mnemonic.lhu       => Format.i
            case Mnemonic.sb        => Format.s
            case Mnemonic.sh        => Format.s
            case Mnemonic.sw        => Format.s
            case Mnemonic.addi      => Format.i
            case Mnemonic.slti      => Format.i
            case Mnemonic.sltiu     => Format.i
            case Mnemonic.xori      => Format.i
            case Mnemonic.ori       => Format.i
            case Mnemonic.andi      => Format.i
            case Mnemonic.slli      => Format.i
            case Mnemonic.srli      => Format.i
            case Mnemonic.srai      => Format.i
            case Mnemonic.add       => Format.r
            case Mnemonic.sub       => Format.r
            case Mnemonic.sll       => Format.r
            case Mnemonic.slt       => Format.r
            case Mnemonic.sltu      => Format.r
            case Mnemonic.xor       => Format.r
            case Mnemonic.srl       => Format.r
            case Mnemonic.sra       => Format.r
            case Mnemonic.or        => Format.r
            case Mnemonic.and       => Format.r
            case Mnemonic.fence     => Format.i
            case Mnemonic.fence_i   => Format.i
            case Mnemonic.ecall     => Format.i
            case Mnemonic.ebreak    => Format.i
            case Mnemonic.csrrw     => Format.i
            case Mnemonic.csrrs     => Format.i
            case Mnemonic.csrrc     => Format.i
            case Mnemonic.csrrwi    => Format.i
            case Mnemonic.csrrsi    => Format.i
            case Mnemonic.csrrci    => Format.i
            //case Mnemonic.invalid   =>
        }

        val instr = fmt match {
            case Format.r =>
                new InstrR(mnem, fmt, rd, func3, rs1, rs2, func7)
            case Format.i =>
                new InstrI(mnem, fmt, rd, func3, rs1, imm_11_0)
            case Format.s =>
                new InstrS(mnem, fmt, imm_4_0, func3, rs1, rs2, imm_11_5)
            case Format.b =>
                new InstrB(mnem, fmt, imm_11_a7, imm_4_1, func3, rs1, rs2, imm_12, imm_10_5)
            case Format.u =>
                new InstrU(mnem, fmt, rd, imm_31_12)
            case Format.j =>
                new InstrJ(mnem, fmt, rd, imm_19_12, imm_11_a20, imm_10_1, imm_20)
        }

        return instr
    }

    def reset(): Unit = {

        for (i <- 0 until regs.length) {
            regs(i) = 0
        }
        pc = 0

        // TODO: remove these; put them in the test itself
        regs(Register.A1.id) = 5    // n
        regs(Register.A0.id) = 0x54 // &a[0]
        regs(Register.RA.id) = 0x4c // return address

        // TODO: randomize memory upon reset?
    }

    def step(): Boolean = {
        var op = M_r(pc, 4)

        val compressed = (op & 3) != 3
        if (compressed)
            op &= 0xffff
        val instr = decode(op)
        var next_pc = pc + (if (compressed) 2 else 4)

        var xd_set = false
        var write = false
        var load = false
        var br_taken = false

        instr.mnem match {
            case Mnemonic.lui =>
                regs(instr.rd) = instr.imm
                xd_set = true
            case Mnemonic.auipc =>
                regs(instr.rd) = pc + instr.imm
                xd_set = true
            case Mnemonic.jal =>
                regs(instr.rd) = pc + 4
                next_pc = pc + instr.imm
                br_taken = true
                xd_set = true
            case Mnemonic.jalr =>
                val t = pc + 4
                next_pc = (regs(instr.rs1) + instr.imm) & ~1
                br_taken = true
                regs(instr.rd) = t
                xd_set = true
            case Mnemonic.beq =>
                if (regs(instr.rs1) == regs(instr.rs2)) {
                    next_pc = pc + instr.imm
                    br_taken = true
                }
            case Mnemonic.bne =>
                if (regs(instr.rs1) != regs(instr.rs2)) {
                    next_pc = pc + instr.imm
                    br_taken = true
                }
            case Mnemonic.blt =>
                if (regs(instr.rs1) < regs(instr.rs2)) {
                    next_pc = pc + instr.imm
                    br_taken = true
                }
            case Mnemonic.bge =>
                if (regs(instr.rs1) >= regs(instr.rs2)) {
                    next_pc = pc + instr.imm
                    br_taken = true
                }
            case Mnemonic.bltu =>
                // TODO: unsigned
                if (regs(instr.rs1) < regs(instr.rs2)) {
                    next_pc = pc + instr.imm
                    br_taken = true
                }
            case Mnemonic.bgeu =>
                // TODO: unsigned
                if (regs(instr.rs1) >= regs(instr.rs2)) {
                    next_pc = pc + instr.imm
                    br_taken = true
                }
            case Mnemonic.lb =>
                regs(instr.rd) = sext(M_r(regs(instr.rs1) + instr.imm, 1), 8)
                xd_set = true
                load = true
            case Mnemonic.lh =>
                regs(instr.rd) = sext(M_r(regs(instr.rs1) + instr.imm, 2), 16)
                xd_set = true
                load = true
            case Mnemonic.lw =>
                regs(instr.rd) = sext(M_r(regs(instr.rs1) + instr.imm, 4), 32)
                xd_set = true
                load = true
            case Mnemonic.lbu =>
                regs(instr.rd) = M_r(regs(instr.rs1) + instr.imm, 1)
                xd_set = true
            case Mnemonic.lhu =>
                regs(instr.rd) = M_r(regs(instr.rs1) + instr.imm, 2)
                xd_set = true
            case Mnemonic.sb =>
                M_w(regs(instr.rs1) + instr.imm, regs(instr.rs2), 1)
                write = true
            case Mnemonic.sh =>
                M_w(regs(instr.rs1) + instr.imm, regs(instr.rs2), 2)
                write = true
            case Mnemonic.sw =>
                M_w(regs(instr.rs1) + instr.imm, regs(instr.rs2), 4)
                write = true
            case Mnemonic.addi =>
                regs(instr.rd) = regs(instr.rs1) + instr.imm
                xd_set = true
            case Mnemonic.slti =>
                regs(instr.rd) = if (regs(instr.rs1) < instr.imm) 1 else 0
                xd_set = true
            case Mnemonic.sltiu =>
                // TODO: unsigned
                regs(instr.rd) = if (regs(instr.rs1) < instr.imm) 1 else 0
                xd_set = true
            case Mnemonic.xori =>
                regs(instr.rd) = regs(instr.rs1) ^ instr.imm
                xd_set = true
            case Mnemonic.ori =>
                regs(instr.rd) = regs(instr.rs1) | instr.imm
                xd_set = true
            case Mnemonic.andi =>
                regs(instr.rd) = regs(instr.rs1) & instr.imm
                xd_set = true
            case Mnemonic.slli =>
                regs(instr.rd) = regs(instr.rs1) << instr.shamt
                xd_set = true
            case Mnemonic.srli =>
                // TODO: unsigned
                regs(instr.rd) = regs(instr.rs1) >> instr.shamt
                xd_set = true
            case Mnemonic.srai =>
                /* TODO: add static assertion that right-shifting of signed value is arithmetic */
                regs(instr.rd) = regs(instr.rs1) >> instr.shamt
                xd_set = true
            case Mnemonic.add =>
                regs(instr.rd) = regs(instr.rs1) + regs(instr.rs2)
                xd_set = true
            case Mnemonic.sub =>
                regs(instr.rd) = regs(instr.rs1) - regs(instr.rs2)
                xd_set = true
            case Mnemonic.sll =>
                regs(instr.rd) = regs(instr.rs1) << regs(instr.rs2)
                xd_set = true
            case Mnemonic.slt =>
                regs(instr.rd) = if (regs(instr.rs1) < regs(instr.rs2)) 1 else 0
                xd_set = true
            case Mnemonic.sltu =>
                // TODO: unsigned
                regs(instr.rd) = if (regs(instr.rs1) < regs(instr.rs2)) 1 else 0
                xd_set = true
            case Mnemonic.xor =>
                regs(instr.rd) = regs(instr.rs1) ^ regs(instr.rs1)
                xd_set = true
            case Mnemonic.srl =>
                // TODO: unsigned
                regs(instr.rd) = regs(instr.rs1) >> regs(instr.rs2)
                xd_set = true
            case Mnemonic.sra =>
                regs(instr.rd) = regs(instr.rs1) >> regs(instr.rs2)
                xd_set = true
            case Mnemonic.or =>
                regs(instr.rd) = regs(instr.rs1) | regs(instr.rs1)
                xd_set = true
            case Mnemonic.and =>
                regs(instr.rd) = regs(instr.rs1) & regs(instr.rs1)
                xd_set = true
            case Mnemonic.fence =>
                /* TODO */
            case Mnemonic.fence_i =>
                /* TODO */
            case Mnemonic.ecall =>
                /* TODO */
            case Mnemonic.ebreak =>
                return true
                /* TODO */
            case Mnemonic.csrrw =>
                /* TODO */
            case Mnemonic.csrrs =>
                /* TODO */
            case Mnemonic.csrrc =>
                /* TODO */
            case Mnemonic.csrrwi =>
                /* TODO */
            case Mnemonic.csrrsi =>
                /* TODO */
            case Mnemonic.csrrci =>
                /* TODO */
        }

        println(f"${pc}%04x: ${instr}")
        pc = next_pc

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
            println(f"M[${i*16}%04x]: ${M_r(i*16, 4)}%08x ${M_r(i*16+4, 4)}%08x ${M_r(i*16+8, 4)}%08x ${M_r(i*16+12, 4)}%08x")
        }
        println("")
    }

    def main(args: Array[String]) = {
        val filename = args(0)
        val ops = ArrayBuffer[Int]()
        for (line <- Source.fromFile(filename).getLines) {
            ops.append(hexToInt(line))
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
