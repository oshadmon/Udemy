/**
 * Enums logic: 
 *      ADMIN = 0 
 *      READ_ONLY = 1 
 *      AUTHOR = 2 
 */
enum Roles {ADMIN, READ_ONLY, AUTHOR}; 

//declare object with no details: this is good when data in object is manipulated
const person:object = { 
    name: 'Ori', 
    age: 30,
    hobbies: ['Sports', 'Cooking'],
    role: [2, 'author'], 
    permission: Roles.ADMIN
}

/**
// declare object with details 
const person:{
    name:string,
    age:number, 
    hobbies: string[],
    role: [number, string], // tuple of number (place 0) & strings (place 1), 

} = { 
    name: 'Ori', 
    age: 30,
    hobbies: ['Sports', 'Cooking'],
    role: [2, 'author'], 
}*/

// print full object 
console.log(person)

// print value for key name
// if you sepcify  type
console.log(person['name'])
//if you don't sepcify a type 
//console.log(person.name)

for (const hobby of person['hobbies']){ 
    // itterate through peprson.hbbies 
    console.log(hobby)
}

console.log(person['role'])
// push value into role <-- allows change lenght of tuple
person['role'].push('adm in')
//person['role'][1] = 10 
console.log(person['role'])

// compare value to enum value
if (person['permission'] === Roles.ADMIN){
    console.log('Admin'); 
}
else{ 
    console.log('Not admin'); 
}