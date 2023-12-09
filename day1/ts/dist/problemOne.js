export default (file) => {
    const isNumeric = (input) => {
        return !isNaN(parseFloat(input)) && isFinite(parseFloat(input));
    };
    const lines = file.split('\n');
    let sum = 0;
    for (let line of lines) {
        let numStr = "";
        for (let char of line) {
            if (isNumeric(char)) {
                numStr += char;
            }
        }
        sum += Number(numStr[0] + numStr[numStr.length - 1]);
    }
    return sum;
};
