import "reflect-metadata";

console.log(Reflect)

// const formatMetadataKey = Symbol("format");

// function format(formatString: string) {
//     return Reflect.metadata(formatMetadataKey, formatString);
// }

// function getFormat(target: any, propertyKey: string) {
//     return Reflect.getMetadata(formatMetadataKey, target, propertyKey);
// }

// class Greeter {
//   @format("Hello, %s")
//   greeting: string;

//   constructor(message: string) {
//       this.greeting = message;
//   }
//   greet() {
//       let formatString = getFormat(this, "greeting");
//       return formatString.replace("%s", this.greeting);
//   }
// }

// let g: Greeter = new Greeter("张三");
// g.greet();