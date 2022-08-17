const assert = std.debug.assert;
const print = std.debug.print;
const std = @import("std");
const native_arch = builtin.cpu.arch;
const builtin = @import("builtin");

pub const Vec = struct {
    x: u32,
    y: u32,
    pub fn format(value: Vec, comptime fmt: []const u8, options: std.fmt.FormatOptions, writer: anytype) !void {
        _ = fmt;
        _ = options;
        try writer.print("Vec of ({d}, {d})", .{ value.x, value.y });
    }
};

const Data = struct {
    data: i32,
};

pub fn main() !void {

    const stdout = std.io.getStdOut().writer();
    try stdout.print("hello, world\n", .{});
    print("sup\n", .{});

    var optional_value: ?[]const u8 = null;
    assert(optional_value == null);

    print("optional 1\ntype type: {s}\ntype: {s}\nvalue: {s}\n", .{
        @typeName(@TypeOf(@TypeOf(optional_value))),
        @typeName(@TypeOf(optional_value)),
        optional_value,
    });

    optional_value = "hi";

    print("optional 1\ntype: {s}\nvalue: {s}\n", .{
        @typeName(@TypeOf(optional_value)),
        optional_value,
    });

    var number_or_error: anyerror!i32 = error.ArgNotFound;

    print("\nerror union 1\ntype: {s}\nvalue: {}\n", .{
        @typeName(@TypeOf(number_or_error)),
        number_or_error
    });

    // primitive types
    const a :i8 = -2;
    const b :u8 = 2;
    const c :isize = 5; // intptr_t
    const d :usize = 5; // uintptr_t
    const e :c_int = 7; // int
    const f :f32 = 0.25; // float
    const g :bool = true; // bool
    //const h :anyopaque = ?;
    const h :type = @TypeOf(g);
    //const i :anyerror!void = error.ArgNotFound;
    const j :comptime_int = 2 + 5;

    _ = a;
    _ = b;
    _ = c;
    _ = d;
    _ = e;
    _ = f;
    _ = g;
    _ = h;
    _ = j;

    var v = Vec{ .x = 10, .y = 20 };
    std.log.info("{}", .{v});

    // Some interesting looking namespaces in std:
    //  ascii
    print("{u}\n", .{std.ascii.toUpper('x')});
    //  atomic
    //  base64
    //  build
    //  compress
    //      std.compress.deflate.decompressor(Allocator, reader: anytype, dictionary: anytype) anytype;
    //      std.compress.gzip
    //      std.compress.zlib
    //  debug
    std.debug.dumpCurrentStackTrace(null);
    //  fmt
    //  fs
    print("{}\n", .{std.fs.cwd()});
    //  heap
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer arena.deinit();
    const allocator = arena.allocator();
    const ptr = try allocator.create(i32);
    print("ptr={*}\n", .{ptr});
    //  http
    //  io
    //  json
    const Foo = struct { a: i32, b: bool };
    const s = "{ \"a\": 15, \"b\": true }";
    var stream = std.json.TokenStream.init(s);
    const parsedData = std.json.parse(Foo, &stream, .{});
    _ = try parsedData;
    //  math
    print("{}\n", .{ std.math.ceil(2.1) });                     // 3.0e+00
    print("{}\n", .{ std.math.floor(2.1) });                    // 2.0e+00
    print("{}\n", .{ std.math.round(2.1) });                    // 2.0e+00
    print("{}\n", .{ std.math.fabs(1.3) });                     // 1.3e+00
    print("{}\n", .{ std.math.clamp(5, 2, 8) });                // 5
    print("{}\n", .{ std.math.cos(std.math.pi) });              // -1.0e+00
    print("{}\n", .{ std.math.sin(std.math.pi) });              // 1.2246467991473532e-16
    print("{}\n", .{ std.math.tan(std.math.pi) });              // -1.2246467991473532e-16
    print("{}\n", .{ std.math.pow(f32, std.math.sqrt2, 2) });   // 1.99999988e+00
    print("{}\n", .{ std.math.sqrt(2.0) });                     // 1.4142135623730951e+00
    print("{}\n", .{ std.math.exp(1.0) });                      // 2.7182818284590455e+00
    print("{}\n", .{ std.math.ln(std.math.e) });                // 1.0e+00
    print("{}\n", .{ std.math.log(f32, 3, 9) });                // 2.0e+00
    print("{}\n", .{ std.math.log10(100) });                    // 2
    print("{}\n", .{ std.math.log2(4) });                       // 2
    print("{}\n", .{ std.math.min(1, 2) });                     // 1
    print("{}\n", .{ std.math.max(1, 2) });                     // 2
    print("{}\n", .{ std.math.mod(i32, 8, 3) });                // 2
    print("{}\n", .{ std.math.shl(i32, 2, 1) });                // 4
    print("{}\n", .{ std.math.shr(i32, 2, 1) });                // 1
    //  mem
    //  net
    //  os
    print("{s}\n", .{ std.os.getenv("PATH")} );
    //  process
    //  rand
    //  simd
    //  sort
    var items = [_]i32 {21, -3, 8, 4, 19, 6, 7};
    std.sort.sort(i32, items[0..], {}, comptime std.sort.asc(i32));
    print("{any}\n", .{items});
    const S = struct {
        fn order_i32(context: void, lhs: i32, rhs: i32) std.math.Order {
            _ = context;
            return std.math.order(lhs, rhs);
        }
    };
    print("{}\n", .{std.sort.binarySearch(i32, 6, items[0..], {}, S.order_i32)});
    //  testing
    //  time
    print("{}\n", .{std.time.milliTimestamp()});
    print("Sleeping for 500ms...\n", .{});
    std.time.sleep(500000000);
    print("done\n", .{});
    //  unicode
    //  wasm

    // Some interesting looking Types in std:
    //  ArrayHashMap
    //  ArrayList
    const List = std.ArrayList(Data);
    var list = List.init(allocator);
    defer list.deinit();
    list.append(Data{.data = 5}) catch unreachable;
    print("{}\n", .{list});
    //  HashMap
    //  StringArrayHashMap
    //  StringHashMap
    var my_hash_map = std.StringHashMap(i32).init(allocator);
    try my_hash_map.put("bob", 5);
    try my_hash_map.put("joe", 5);
    print("{}\n", .{my_hash_map.get("bob")});

    // assembly

    var my_asm : [*]usize = undefined;

    switch (native_arch) {
        .riscv64 => {
            my_asm = asm volatile (
                \\ li s0, 0
                \\ li ra, 0
                : [argc] "={sp}" (-> [*]usize),
            );
        },
        //.x86_64 => {
        //    my_asm = asm volatile (
        //        \\ xor %%rbp, %%rbp
        //        : [argc] "={rsp}" (-> [*]usize),
        //    );
        //},
        //else => @compileError("unsupported arch"),
        else => {},
    }
}
