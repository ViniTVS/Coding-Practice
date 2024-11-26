<?php


/**
 * @param string[] $words
 * @param int $maxWidth
 * @return string[]
 */
function fullJustify($words, $maxWidth): array {
  $count = 0;
  $line = [];
  $res = [];
  $words_count = 0;
  $real_count = 0;
  foreach ($words as $word) {
    $w_len = strlen($word);

    if ($count + $w_len <= $maxWidth) {
      $line[] = $word;
      $line[] = "";
      $count += $w_len + 1;
      $real_count += $w_len;
      $words_count++;
    } else {
      $num = $maxWidth - $real_count;
      $words_count = count($line) / 2;
      $denum = max($words_count - 1, 1);
      $div = intval(floor($num / $denum));
      $mod = $num % $denum;
      // var_dump("div: $div, mod: $mod, word_count: $words_count, aux: $aux");
      for ($j = 0; $j < max($words_count - 1, 1); $j++) {
        // var_dump("line: ", 1 + $j * 2, $line);
        $line[1 + $j * 2] = $j < $mod ? str_repeat(" ", $div + 1) : str_repeat(" ", $div);
      }
      $res[] = implode("", $line);
      $line = [0 => $word];
      $line[] = "";
      $words_count = 1;
      $real_count = $w_len;
      $count = $w_len + 1;
    }
  }
  if (count($line)) {
    if ($real_count == $maxWidth) {
    $res[] = implode("", $line);
    }
    for ($j = 0; $j < max($words_count - 1, 1); $j++) {
      // var_dump("line: ", 1 + $j * 2, $line);
      $line[1 + $j * 2] = " ";
    }
    $aux = implode("", $line);
    $end = str_repeat(" ", $maxWidth - strlen($aux));
    $res[] = $aux . $end;
  }
  return $res;
}

var_dump(fullJustify(["What","must","be","acknowledgment","shall","be"], 16));
// ["This    is    an","example  of  text","justification.  "]
// ["This    is    an","example  of text","justification.  "]