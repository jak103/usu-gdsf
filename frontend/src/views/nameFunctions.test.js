import stringToObjects from "./nameFunctions";

let strings = [
  "firstname1 lastname1",
  "firstname2 lastname2"
];

let objects = [
  {
    id: 0,
    firstName: "firstname1",
    lastName: "lastname1"
  },
  {
    id: 1,
    firstName: "firstname2",
    lastName: "lastname2"
  }
];

test('Empty list', () => {
  expect(stringToObjects([])).toStrictEqual([]);
});

test('Valid list', () => {
  strings.forEach((fullName, index) => {
    let firstName = fullName.split(" ")[0];
    let lastName = fullName.split(" ")[1];
    expect(index).toBe(objects[index].id);
    expect(firstName).toBe(objects[index].firstName);
    expect(lastName).toBe(objects[index].lastName);
  });
});

  