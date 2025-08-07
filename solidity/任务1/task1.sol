/*题目1-创建一个名为Voting的合约 开始*/
mapping(string name => uint32 amount) public voteMapping;

    struct Hobbie {
        string label;
        string value; 
    }

    string[] private names;

    function vote(string calldata name) public {
        if(voteMapping[name] == 0 && bytes(name).length > 0){
            names.push(name);
        }
        voteMapping[name]++;
    }

    function getVotes(string calldata name) public view returns (uint32) {
        return voteMapping[name];
    }

    function resetVotes() public  {
        for(uint256 i = 0; i < names.length; i++){
            voteMapping[names[i]] = 0;
        }
        
        delete names;
    }
/*题目1-创建一个名为Voting的合约 结束*/


/*题目2-反转字符串 开始*/
 function reverse(string calldata str) public pure returns (string memory) {
        bytes memory bStr  = bytes(str);
        bytes memory reverseStr = new bytes(bStr.length);
        uint256 size = bStr.length;
        for(uint256 i = 0; i < size; i++){
            reverseStr[i] = bStr[bStr.length-1-i];
        }
        return string(reverseStr);
    
    }
/*题目2-反转字符串 结束*/

/*题目3-用 实现整数转罗马数字 开始*/
//将数字转换为罗马字符串 1348   23
    function IntToRoman(uint256 i256) public pure returns (string memory) {
        require(i256 > 0 && i256 < 4000, "Number must be 1-3999");
        
        string memory result;
        uint256 temp = i256;
        
        // 千位处理
        while (temp >= 1000) {
            result = string.concat(result, "M");
            temp -= 1000;
        }
        
        // 百位处理
        if (temp >= 900) {
            result = string.concat(result, "CM");
            temp -= 900;
        } else if (temp >= 500) {
            result = string.concat(result, "D");
            temp -= 500;
            while (temp >= 100) {
                result = string.concat(result, "C");
                temp -= 100;
            }
        } else if (temp >= 400) {
            result = string.concat(result, "CD");
            temp -= 400;
        } else while (temp >= 100) {
            result = string.concat(result, "C");
            temp -= 100;
        }
        
        // 十位处理
        if (temp >= 90) {
            result = string.concat(result, "XC");
            temp -= 90;
        } else if (temp >= 50) {
            result = string.concat(result, "L");
            temp -= 50;
            while (temp >= 10) {
                result = string.concat(result, "X");
                temp -= 10;
            }
        } else if (temp >= 40) {
            result = string.concat(result, "XL");
            temp -= 40;
        } else while (temp >= 10) {
            result = string.concat(result, "X");
            temp -= 10;
        }
        
        // 个位处理
        if (temp >= 9) {
            result = string.concat(result, "IX");
            temp -= 9;
        } else if (temp >= 5) {
            result = string.concat(result, "V");
            temp -= 5;
            while (temp >= 1) {
                result = string.concat(result, "I");
                temp -= 1;
            }
        } else if (temp >= 4) {
            result = string.concat(result, "IV");
            temp -= 4;
        } else while (temp >= 1) {
            result = string.concat(result, "I");
            temp -= 1;
        }
        
        return result;
    }
/*题目3-用 实现整数转罗马数字 结束*/


/*题目4-用 实现罗马数字转数整数 开始*/
pragma solidity ^0.8.0;
contract RomanToInteger {
   //映射罗马对应的数字
   mapping(bytes1 => uint256) private _numericValues;

   constructor(){
     _numericValues['I'] = 1;
     _numericValues['V'] = 5;
     _numericValues['X'] = 10;
     _numericValues['L'] = 50;
     _numericValues['C'] = 100;
     _numericValues['D'] = 500;
     _numericValues['M'] = 1000;
   }

   //将罗马字符串转换为数字
   function romanToInt(string memory s) public view returns (uint256){
        //结果
        uint256 result = 0;
        uint256 before = 0;
        bytes memory bytesS = bytes(s);

         for(int256 i = int256(bytesS.length-1); i >= 0; i--){
            bytes1 currentKey = bytesS[uint(i)];
            uint256 currentVlue = _numericValues[currentKey];
             //如果当前位大于 前一位,则加当前位, 如果当前位小于前一位, 则减当前位 
             if(currentVlue >= before){
                  result += currentVlue; 
             }else{
                  result -= currentVlue;
             }
             before = currentVlue;
         }
         return result;
   }
}
/*题目4-用 实现罗马数字转数整数 结束*/


/*题目5- 合并两个有序数组 开始*/
contract RomanToInteger {
    
    function romanToInt(uint256[] memory array1, uint256[] memory array2) public pure returns (uint256[] memory){
        uint256 length1 = array1.length;
        uint256 length2 = array2.length;

        uint256 length3 = length1 + length2;

         uint256[] memory array3 = new uint256[](length3);
        uint256 i = 0;
        uint256 j = 0;
        uint256 k = 0;

        while(i < length1 && j < length2){
            if(array1[i] < array2[j]){
                array3[k] = array1[i];
                i++;
            }else{
                array3[k] = array2[j];
                j++;
            }
            k++;
        }

        if(i == length1){
            for(uint m = j; m < length2; m++){
                array3[k] = array2[m];
                k++;
            }
        }
        if(j == length1){
            for(uint m = i; m < length1; m++){
                array3[k] = array2[m];
                k++;
            }
        }
        return array3;
    
    }
}
/*题目5- 合并两个有序数组 结束*/


/*题目6- 二分查找 开始*/
contract BinarySearch {
    function romanToInt(uint256[] memory array, uint256 value) public pure returns (int256){
        uint256 start = 0;
        uint256 end = array.length - 1 ;
        uint256 middleIndex = (start+end)/2;
        while(start < end){
            if(array[middleIndex] < value) {
                //go to right half 
                start = middleIndex + 1;
            }else{
                end = middleIndex;
            }
            middleIndex = (start+end)/2;
        }

        if(array[middleIndex] == value){
            return int256(middleIndex);
        }else{
            return -1;
        }

    }
}
/*题目6- 二分查找 结束*/
