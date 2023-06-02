
async function nlpURL(url) {
   try {
      const response = await fetch(
         `http://localhost:8081/nlp?url=${encodeURIComponent(url)}`
      );

      if (!response.ok) {
         throw new Error(`HTTP error: ${response.status}`);
      }

      const data = await response.json();
      return data;
   } catch (error) {
      throw new Error(`HTTP error: ${error}`);
   }
}

export { nlpURL };
