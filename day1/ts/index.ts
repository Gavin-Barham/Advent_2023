import * as fs from 'fs';
import problemOne from './problemOne.js';
import problemTwo from './problemTwo.js';

interface data {
    input: string;
    test: string;
    testTwo: string;
}
const getData = (): data => {
    const input  = fs.readFileSync('./input.txt', 'utf8');
    const test  = fs.readFileSync('./test.txt', 'utf8');
    const testTwo  = fs.readFileSync('./test2.txt', 'utf8');


    return {input, test, testTwo};
}
const {input, test, testTwo} = getData();

// const test1 = problemOne(test);
// const answer1 = problemOne(input);
// console.log(`problem one test:{ passed:(${test1 === 142}), value:(${test1})}`);
// console.log(`problem one answer:{ passed:(${answer1 === 54630}), value:(${answer1})}`);
const test2 = problemTwo(testTwo);
const answer2 = problemTwo(input);
console.log(`problem two test:{ passed:(${test2 === 281}), value:(${test2})}`);
console.log(`problem two answer:{ passed:(${answer2 === 54770}), value:(${answer2})}`);

