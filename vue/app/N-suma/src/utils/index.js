export { roles, clases, requiredRules, numberRules, divisions };
const roles = ["学生", "教員"];
const clases = [
  "CTA20",
  "CTB20",
  "CTA21",
  "CTB21",
  "CTA22",
  "CTB22",
  "CTA23",
  "CTB23",
];
const requiredRules = [(v) => !!v || "必須"];
const numberRules = [
  (v) => !!v || "必須",
  (v) => v.length == 8 || "学籍番号は8桁です",
  (v) => /^\d+$/.test(v) || "学籍番号は半角数字です",
];
const divisions = ["国家試験 / FE", "国家試験 / AP"];
