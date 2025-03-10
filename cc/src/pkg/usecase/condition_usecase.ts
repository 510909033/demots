interface Animal {
    live(): void;
    v2(): void;
  }
  interface Dog extends Animal {
    woof(): void;
  }
  
//   type Example1 = Dog extends Animal ? number : string;
  
//   type Example2 = RegExp extends Animal ? number : string;

export class DemoCondition {
    demo1() {
        type Example1 = Dog extends Animal ? number : string;
        type Example2 = RegExp extends Animal ? number : string;

        let e1: Example1 = 1;
    }
}


// interface IdLabel {
//     id: number /* some fields */;
//   }
//   interface NameLabel {
//     name: string /* other fields */;
//   }
  
//   function createLabel(id: number): IdLabel;
//   function createLabel(name: string): NameLabel;
//   function createLabel(nameOrId: string | number): IdLabel | NameLabel;
//   function createLabel(nameOrId: string | number): IdLabel | NameLabel {
//     throw "unimplemented";
//   }

//   type NameOrId<T extends number | string> = T extends number
//   ? IdLabel
//   : NameLabel;


  interface IdLabel {
    id: number /* some fields */;
  }
  interface NameLabel {
    name: string /* other fields */;
  }
  type NameOrId<T extends number | string> = T extends number
    ? IdLabel
    : NameLabel;
  function createLabel<T extends number | string>(idOrName: T): NameOrId<T> {
    if (typeof idOrName === "string") {
      return { name: idOrName } as NameOrId<T>;
    }
    return { id: idOrName } as NameOrId<T>;
  }
  
//   type IntegerOrString<T extends number | string> = T extends number ? Integer<T> : T;
//   let i2:IntegerOrString<number> = 1.2;
//   console.log("-----------",i2);