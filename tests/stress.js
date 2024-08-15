import http from "k6/http";

export default function () {
  http.get("http://localhost:30202/customers?cpf=12345678910");

}
