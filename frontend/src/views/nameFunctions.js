export default function stringToObjects(stringList) {
  let objectList = [];
  let i = 0;
  while (i !== stringList.length) {
    let fullName = stringList[i].split(" ");
    objectList[i] = {id: i, firstName: fullName[0], lastName: fullName[1]};
    i++;
  }
  return objectList;
}