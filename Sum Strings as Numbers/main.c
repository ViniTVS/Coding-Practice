#include <stdlib.h>
#include <string.h>
#include <stdio.h>

// This shouldn't be a 4 kuy kata... 
char *strsum(const char *a, const char *b) {

    //  <----  hajime!
    // declare variables...
    int len_a, len_b, max, index, val, carry_one = 0;
    char * final_str;
    char c;

    len_a = strlen(a);
    len_b = strlen(b);
    max = (len_a > len_b) ? len_a : len_b;
    final_str = calloc(max + 2, sizeof(char)); 

    index = max + 1;
    // add numbers just like you learned at school
    for( int i = 1; i <= max; i++){

        val = carry_one;
        // chr to int
        if (len_a - i >= 0)
            val += a[len_a - i] - '0';
        if (len_b - i >= 0)
            val += b[len_b - i] - '0';
        
        carry_one = val / 10; // carry one to next digit?
        val = val % 10;
        c = val + '0'; // int to chr
        index--;
        final_str[index] = c;
    }
    
    final_str[max + 1] = '\0'; // str limiter
    // check if there's one more digit
    if (carry_one){
        index--;
        final_str[index] = carry_one + '0';
    }
    // count the number of zeros on the final string 
    int zeros = 0;
    while(final_str[zeros] == '0' || final_str[zeros] == 0){
        zeros++;
    }
    
    if (zeros){
        // the result is 0
        if (max + 2 - zeros < 2){
            free(final_str);
            char *str = calloc(2, sizeof(char));
            str[0] = '0';
            return str;
        }
        // remove unnecessary 0s
        char *aux_str = calloc(max + 2 - zeros, sizeof(char)); 
        memcpy(aux_str, &(final_str[zeros]), sizeof(char) * (max + 2 - zeros));
        free(final_str);
        return aux_str;
    }
    
    return final_str;
}

int main(){
    char *teste = strsum("00000", "0");
    printf("saida: %s \n", teste);
    return 0;
}