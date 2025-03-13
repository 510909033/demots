enum Status{
    OK = 200,
    CREATED = 201,
    ACCEPTED = 202,
    NO_CONTENT = 204,
    BAD_REQUEST = 400,
    UNAUTHORIZED = 401,
    FORBIDDEN = 403,
    NOT_FOUND = 404,
    METHOD_NOT_ALLOWED = 405,
    CONFLICT = 409,
    UNPROCESSABLE_ENTITY = 422,
}


enum Direction {
    Up = 1,
    Down,
    Left,
    Right,
  }

const getSomeValue = () => 23;
enum E {
  A = getSomeValue(),
  B = 23,
}

console.log(E)


enum ShapeKind {
    Circle,
    Square,
  }
  
  interface Circle {
    kind: ShapeKind.Circle;
    radius: number;
  }
  
  interface Square {
    kind: ShapeKind.Square;
    sideLength: number;
  }
  
  let c: Circle = {
    kind: ShapeKind.Circle,
    radius: 100,
  };

export default Status;
export {Direction};
enum E {
    X,
    Y,
    Z,
  }
  
  function f(obj: { X: number }) {
    return obj.X;
  }
  
  // Works, since 'E' has a property named 'X' which is a number.
  f(E);
  
//   let ff:E = { X: 0, Y: 1}