import { checkForURL, checkForName } from "../src/js/nameChecker"

describe("Testing nameCheck functionality", () => {
   test("Testing nlpURL() function", () => {
      expect(checkForURL).toBeDefined();
      expect(checkForURL("foobar")).toBe(false);
      expect(checkForURL("https://www.google.com")).toBe(true);
   }),
   test("Testing checkForName", () => {
      expect(checkForName).toBeDefined();
      expect(checkForName("foobar")).toBeNull();

   })
});

