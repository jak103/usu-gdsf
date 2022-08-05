import stringsToObjects from "./nameFunctions";

let strings = [
  "firstname1 lastname1",
  "firstname2 lastname2"
];

let objectsExpected = [
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
  expect(stringsToObjects([])).toStrictEqual([]);
});

test('Valid list', () => {
  let objects = stringsToObjects(strings);
  objects.forEach((object, index) => {
    expect(index).toBe(objectsExpected[index].id);
    expect(object.firstName).toBe(objectsExpected[index].firstName);
    expect(object.lastName).toBe(objectsExpected[index].lastName);
  });
});

  