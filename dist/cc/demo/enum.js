"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.Direction = void 0;
var Status;
(function (Status) {
    Status[Status["OK"] = 200] = "OK";
    Status[Status["CREATED"] = 201] = "CREATED";
    Status[Status["ACCEPTED"] = 202] = "ACCEPTED";
    Status[Status["NO_CONTENT"] = 204] = "NO_CONTENT";
    Status[Status["BAD_REQUEST"] = 400] = "BAD_REQUEST";
    Status[Status["UNAUTHORIZED"] = 401] = "UNAUTHORIZED";
    Status[Status["FORBIDDEN"] = 403] = "FORBIDDEN";
    Status[Status["NOT_FOUND"] = 404] = "NOT_FOUND";
    Status[Status["METHOD_NOT_ALLOWED"] = 405] = "METHOD_NOT_ALLOWED";
    Status[Status["CONFLICT"] = 409] = "CONFLICT";
    Status[Status["UNPROCESSABLE_ENTITY"] = 422] = "UNPROCESSABLE_ENTITY";
})(Status || (Status = {}));
var Direction;
(function (Direction) {
    Direction[Direction["Up"] = 1] = "Up";
    Direction[Direction["Down"] = 2] = "Down";
    Direction[Direction["Left"] = 3] = "Left";
    Direction[Direction["Right"] = 4] = "Right";
})(Direction || (exports.Direction = Direction = {}));
const getSomeValue = () => 23;
var E;
(function (E) {
    E[E["A"] = getSomeValue()] = "A";
    E[E["B"] = 23] = "B";
})(E || (E = {}));
console.log(E);
var ShapeKind;
(function (ShapeKind) {
    ShapeKind[ShapeKind["Circle"] = 0] = "Circle";
    ShapeKind[ShapeKind["Square"] = 1] = "Square";
})(ShapeKind || (ShapeKind = {}));
let c = {
    kind: ShapeKind.Circle,
    radius: 100,
};
exports.default = Status;
(function (E) {
    E[E["X"] = 0] = "X";
    E[E["Y"] = 1] = "Y";
    E[E["Z"] = 2] = "Z";
})(E || (E = {}));
function f(obj) {
    return obj.X;
}
// Works, since 'E' has a property named 'X' which is a number.
f(E);
//   let ff:E = { X: 0, Y: 1}
