enum Datatype {
  "STRING" = "STRING",
  "NUMBER" = "NUMBER",
  "BOOLEAN" = "BOOLEAN",
  "TRIGGER" = "TRIGGER",
  "ANY" = "ANY",
  "UI" = "UI",
}

enum DatatypeColors {
  "STRING" = "#DB3A2E",
  "NUMBER" = "#328CDB",
  "BOOLEAN" = "#40DB4F",
  "TRIGGER" = "#787878",
  "ANY" = "#000000",
  "UI" = "#FF8ADC",
}

class Datatypes {
  static allowedConnection(type1: Datatype, type2: Datatype): boolean {
    if (type1 == type2) return true;
    if (type2 == Datatype.ANY && type1 != Datatype.TRIGGER && type1 != Datatype.UI) return true;

    return false;
  }
}

export { Datatype, DatatypeColors, Datatypes };
