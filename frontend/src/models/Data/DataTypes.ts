enum Datatype {
  "string" = "string",
  "number" = "number",
  "bool" = "bool",
  "trigger" = "trigger",
  "any" = "any",
  "ui" = "ui",
}

enum DatatypeColors {
  "string" = "#DB3A2E",
  "number" = "#328CDB",
  "bool" = "#40DB4F",
  "trigger" = "#787878",
  "any" = "#000000",
  "ui" = "#FF8ADC",
}

class Datatypes {
  static allowedConnection(type1: Datatype, type2: Datatype): boolean {
    console.log(type1, type2);
    console.log(type1 == Datatype.string, type2 == Datatype.any);
    if (type1 == type2) return true;
    if (type2 == Datatype.any && type1 != Datatype.trigger && type1 != Datatype.ui) return true;

    return false;
  }
}

export { Datatype, DatatypeColors, Datatypes };
