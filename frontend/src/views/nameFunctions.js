export default function stringsToObjects(stringList) {
  let objectList = [];
  let i = 0;
  while (i !== stringList.length) {
    let fullName = stringList[i].split(" ");
    objectList[i] = {id: i, firstName: fullName[0], lastName: fullName[1]};
    i++;
  }
  return sortObjectListByLastname(objectList);
}

function sortObjectListByLastname(objectList) {
  return objectList.sort((n1, n2) => n1.lastName > n2.lastName ? 1 : -1);
}
