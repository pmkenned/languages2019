"use strict";

// [Understanding Functions and 'this' In The World of ES2017](https://www.youtube.com/watch?v=AOSYY1_np_4)

function Tom(foo) {
    this.foo = foo;
}

function f(x) {
    console.log(`f, x: ${x} this: `, this);
}

console.assert(typeof null                      === "object");
console.assert(typeof undefined                 === "undefined");
console.assert(typeof 5                         === "number");
console.assert(typeof 10n                       === "bigint");
console.assert(typeof "hello"                   === "string");
console.assert(typeof true                      === "boolean");
console.assert(typeof Symbol("x")               === "symbol");
console.assert(typeof Math                      === "object");
console.assert(typeof []                        === "object");
console.assert(typeof {}                        === "object");
console.assert(typeof (() => {})                === "function");
console.assert(typeof Array                     === "function");
console.assert(typeof Object                    === "function");
console.assert(typeof Function                  === "function");
console.assert(typeof String                    === "function");
console.assert(typeof Boolean                   === "function");
console.assert(typeof Number                    === "function");
console.assert(typeof BigInt                    === "function");
console.assert(typeof Window                    === "function");
console.assert(typeof window                    === "object");
console.assert(typeof new Array                 === "object");
console.assert(typeof new Object                === "object");
console.assert(typeof new Function              === "function");
console.assert(typeof new String                === "object");
console.assert(typeof new Boolean               === "object");
console.assert(typeof new Number                === "object");
console.assert(typeof BigInt(1)                 === "bigint");
console.assert(typeof Array.prototype           === "object");
console.assert(typeof Object.prototype          === "object");
console.assert(typeof Function.prototype        === "function");
console.assert(typeof String.prototype          === "object");
console.assert(typeof Boolean.prototype         === "object");
console.assert(typeof Number.prototype          === "object");
console.assert(typeof BigInt.prototype          === "object");
console.assert(typeof f                         === "function");
console.assert(typeof f.prototype               === "object");

// X instanceof Y if Y is in the prototype chain of X

console.assert(Object.getPrototypeOf({})            === Object.prototype);
console.assert(Object.getPrototypeOf([])            === Array.prototype);
console.assert(Object.getPrototypeOf(()=>{})        === Function.prototype);
console.assert(Object.getPrototypeOf("")            === String.prototype);
console.assert(Object.getPrototypeOf(true)          === Boolean.prototype);
console.assert(Object.getPrototypeOf(1)             === Number.prototype);
console.assert(Object.getPrototypeOf(1n)            === BigInt.prototype);
console.assert(Object.getPrototypeOf(window)        === Window.prototype);

console.assert({}           instanceof Object);
console.assert([]           instanceof Array);
console.assert((() => {})   instanceof Function);
console.assert(!(""         instanceof String)); // why?
console.assert(!(true       instanceof Boolean)); // why?
console.assert(!(1          instanceof Number)); // why?
console.assert(!(1n         instanceof BigInt)); // why?
console.assert(window       instanceof Window);

console.assert(Object.getPrototypeOf(new Object)    === Object.prototype);
console.assert(Object.getPrototypeOf(new Array)     === Array.prototype);
console.assert(Object.getPrototypeOf(new Function)  === Function.prototype);
console.assert(Object.getPrototypeOf(new String)    === String.prototype);
console.assert(Object.getPrototypeOf(new Boolean)   === Boolean.prototype);
console.assert(Object.getPrototypeOf(new Number)    === Number.prototype);
console.assert(Object.getPrototypeOf(BigInt(1n))    === BigInt.prototype); // cannot use new with BigInt

console.assert(new Object   instanceof Object);
console.assert(new Array    instanceof Array);
console.assert(new Function instanceof Function);
console.assert(new String   instanceof String);
console.assert(new Boolean  instanceof Boolean);
console.assert(new Number   instanceof Number);
console.assert(!(BigInt(1)  instanceof BigInt)); // why?

console.assert(Object.getPrototypeOf(Object)        === Function.prototype);
console.assert(Object.getPrototypeOf(Array)         === Function.prototype);
console.assert(Object.getPrototypeOf(Function)      === Function.prototype);
console.assert(Object.getPrototypeOf(String)        === Function.prototype);
console.assert(Object.getPrototypeOf(Boolean)       === Function.prototype);
console.assert(Object.getPrototypeOf(Number)        === Function.prototype);
console.assert(Object.getPrototypeOf(BigInt)        === Function.prototype);
console.assert(!(Object.getPrototypeOf(Window)      === Function.prototype)); // why?

console.assert(Object       instanceof Function);
console.assert(Array        instanceof Function);
console.assert(Function     instanceof Function);
console.assert(String       instanceof Function);
console.assert(Number       instanceof Function);
console.assert(BigInt       instanceof Function);
console.assert(Window       instanceof Function);

console.assert(Object       instanceof Object);
console.assert(Array        instanceof Object);
console.assert(Function     instanceof Object);
console.assert(String       instanceof Object);
console.assert(Number       instanceof Object);
console.assert(BigInt       instanceof Object);
console.assert(Window       instanceof Object);


console.assert(Object.getPrototypeOf(Object.prototype)      === null);
console.assert(Object.getPrototypeOf(String.prototype)      === Object.prototype);
console.assert(Object.getPrototypeOf(Number.prototype)      === Object.prototype);
console.assert(Object.getPrototypeOf(Array.prototype)       === Object.prototype);
console.assert(Object.getPrototypeOf(BigInt.prototype)      === Object.prototype);
console.assert(Object.getPrototypeOf(Function.prototype)    === Object.prototype);
//console.assert(Object.getPrototypeOf(Window.prototype)      === Object.prototype); // ??

//console.assert(Object.prototype     instanceof Object); // because Object.getPrototypeOf(Object.prototype) is null
console.assert(Array.prototype      instanceof Object);
console.assert(Function.prototype   instanceof Object);
console.assert(String.prototype     instanceof Object);
console.assert(Boolean.prototype    instanceof Object);
console.assert(Number.prototype     instanceof Object);

console.assert([].prototype         === undefined);
console.assert({}.prototype         === undefined);
console.assert("".prototype         === undefined);
console.assert((()=>{}).prototype   === undefined);
console.assert(window.prototype     === undefined);

console.assert({}.constructor                       === Object);
console.assert([].constructor                       === Array);
console.assert((() => {}).constructor               === Function);
console.assert("".constructor                       === String);
//console.assert(1.constructor                        === Number);
console.assert(1n.constructor                       === BigInt);
console.assert(window.constructor                   === Window);

console.assert(f.length === 1);
console.assert(f.name === "f");
console.assert(f            instanceof Function);
console.assert(f.prototype  instanceof Object);
console.assert(f.prototype.constructor              === f);
console.assert(Object.getPrototypeOf(f)             === Function.prototype);
console.assert(Object.getPrototypeOf(f.prototype)   === Object.prototype);

// ================================================================

console.log(window);
console.log(Window);

//window.f(1);
//f(1.1);
//const f2 = window.f;
//f2(2);
//new f(3);
//f.call({y: 5}, 4);
//const f3 = f.bind({y:6});
//f3(5);

//console.log(Tom);

//f.call(window, 1);
//f.call({y: 3}, 1); // same as new f();
//
//console.assert(this === window);
//
//console.log(typeof Tom.prototype);
//console.log(Tom.prototype.constructor);
//
//console.log(Object);
//console.log(Array);
//console.log(Function);
//
//console.log(window);
//
//const jane = new Tom(2);
////console.assert(Object.getPrototypeOf(jane) === Tom.prototype);
////console.assert(jane instanceof Tom);
////console.log("jane.__proto__: ", jane.__proto__);
////console.assert(Tom.length === 2);
//
////console.assert(Object.getPrototypeOf(Tom) === Function.prototype);
////console.log(Object.prototype);
////console.log(Function.prototype);
////
////Tom.prototype.baz = 1;
////
////console.log(jane);
////console.log(Tom);
////console.log(Tom.prototype.constructor);
////console.log(Object);
//
////console.log(Tom.prototype);
////console.log(Object.getPrototypeOf(jane));
//
//if (false) {
//
//console.log(typeof Object.getPrototypeOf({}));
//console.assert(Object   instanceof      Object);
//console.assert({}       instanceof      Object);
//console.assert(typeof Tom.prototype === "object");
//const jane = new Tom(2);
//
//    console.log(jane);
//    console.log(typeof jane);
//    console.log(typeof Tom);
//    //console.log(typeof function);
//    console.log(Object.getPrototypeOf(Tom));
//    console.log(Object.getPrototypeOf(jane));
//    console.log(Object.getPrototypeOf({}));
//    console.log(Tom.prototype);
//
//    //console.log(document);
//    //const x = "hello";
//    //const y = "hello";
//    //console.log(typeof x);
//    //console.log(x === y);
//    //console.log(x.prototype);
//}
//
// TODO:
// clone
