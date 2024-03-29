import { useStore } from "@/stores/user";
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
const APICallonGET = async (url) => {
  const res = await fetch(import.meta.env.VITE_API_URL + url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  });
  const json = await res.json();
  return json;
};
const APICallonJWT = async (url, data) => {
  const store = useStore();
  const res = await fetch(import.meta.env.VITE_API_URL + url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${store.token}`,
    },
    body: JSON.stringify(data),
  });
  const json = await res.json();
  return json;
};

export {
  roles,
  clases,
  requiredRules,
  numberRules,
  divisions,
  APICallonGET,
  APICallonJWT,
};
