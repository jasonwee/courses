import { nlpURL } from "../src/js/formHandler" 

describe("Testing formHandler functionality", () => {
   test("Testing nlpURL() function", () => {
     expect(nlpURL).toBeDefined();
     nlpURL("https://jasonwee.github.io/qj-srv250/").then(data => {
        expect(data).not.toBeNull();
     });
   })
});

