export default (file: string):number => {
    const numMap = (str: string): string => {
        let result: string = "";
        switch(str) {
            case 'one':
                result = '1'
                break;
            case 'two':
                result = '2'
                break;
            case 'three':
                result = '3'
                break;
            case 'four':
                result = '4'
                break;
            case 'five':
                result = '5'
                break;
            case 'six':
                result = '6'
                break;
            case 'seven':
                result = '7'
                break;
            case 'eight':
                result = '8'
                break;
            case 'nine':
                result = '9'
                break;
            default:
                result = ""   
        }
        return result
    }

    const isNumeric = (input: string): boolean => {
        return !isNaN(parseFloat(input)) && isFinite(parseFloat(input));
    }

    const convertWordToNumeric = (str: string): string => {
        let numStr = "" 
        let numericWord = ""
        for (let i = 0; i < str.length; i++) {
            if (isNumeric(str[i])) {
                numStr += str[i]
                continue
            }
            numericWord = str[i]
            for (let j = i + 1; j < str.length; j++) {
                if (isNumeric(str[j])) {
                    numericWord = ""
                    break
                }
                numericWord += str[j]
                if (numMap(numericWord) !== "") {
                    numStr += numMap(numericWord)
                    numericWord = ''
                    break
                }
            }
        }
        return numStr;
    }

    const lines = file.split('\n');
    let sum = 0;
    for (let line of lines) {
        let numStr = convertWordToNumeric(line);
        sum += Number(numStr[0] + numStr[numStr.length - 1])
    }

    return sum
}

