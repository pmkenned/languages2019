function makeRooter(a) {
    return function (b) {
        return Math.pow(b, 1/a);
    };
}

const sqrt = makeRooter(2);
console.log(sqrt(2));
const cubeRoot = makeRooter(3);
console.log(cubeRoot(8));
