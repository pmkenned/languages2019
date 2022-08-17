const std = @import("std");

pub fn main() !void {
    const x :i32 = -5;
    const y = @as(u32, 5);
    const z :i32 = undefined;
    std.debug.print("x: {} y: {} z: {}\n", .{x, y, z});

    const stdout = std.io.getStdOut().writer();

    try stdout.print("hello!\n", .{});

    const a = [5]u8{ 'h', 'e', 'l', 'l', 'o' };
    const b = [_]u8{ 'w', 'o', 'r', 'l', 'd', '!' };
    const length = a.len;

    try stdout.print("{s} {s} {}\n", .{a, b, length});
}
