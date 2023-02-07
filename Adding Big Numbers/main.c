#include <stdlib.h>
#include <string.h>
#include <stdio.h>

char *add(const char *a, const char *b) {

    //  <----  hajime!
    // declare variables...
    int len_a, len_b, max, index, val, carry_one = 0;
    char * final_str;
    char c;

    len_a = strlen(a);
    len_b = strlen(b);
    max = (len_a > len_b) ? len_a : len_b;
    final_str = calloc(max + 2, sizeof(char)); 

    
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
        index = max + 1 - i;
        final_str[index] = c;
    }
    
    // did not use max + 1 chars to create new string
    // so we can just create a new one 1 char shorter
    if (index > 0){
        if (carry_one == 0){
            char *aux_str;
            aux_str = calloc(max + 1, sizeof(char)); 
            memcpy(aux_str, &(final_str[1]), sizeof(char) * (max + 1));
            free(final_str);
            aux_str[max] = '\0'; // str limiter
            return aux_str;
        } else {
            final_str[0] = '1';
        }
    }
    final_str[max + 1] = '\0'; // str limiter
    return final_str;
}

int main(){
    char *teste = add("888", "222");
    printf("%s \n", teste);
    return 0;
}