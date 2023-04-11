def snail(array)
  # enjoy
  len = array.size
  snail = []
  return [] if array[0].empty?

  index = -1
  for margin in 0..len / 2 - 1
    # top row
    for i in margin..len - 1 - margin
      snail[index += 1] = array[margin][i]
    end
    # right col
    for i in margin + 1..len - 2 - margin
      snail[index += 1] = array[i][len - 1 - margin]
    end
    # bot row
    for i in margin..len - 1 - margin
      snail[index += 1] = array[len - 1 - margin][len - 1 - i]
    end
    # left col
    for i in margin..len - 3 - margin
      snail[index += 1] = array[len - 2 - i][margin]
    end
  end
  # is missing the center value
  snail[index += 1] = array[len / 2][len / 2] if index < len * len - 1
  snail
end

def test(input, _expected)
  output = snail(input)
  for i in output
    print i
    print ' '
  end
  puts ''
  # Test.expect(expected == output, "When snail(#{input}) expected #{expected} but got #{output}")
end
test([[1, 2, 3], [4, 5, 6], [7, 8, 9]], [1, 2, 3, 6, 9, 8, 7, 4, 5])
test(
  [
    [1, 2, 3, 4],
    [5, 6, 7, 8],
    [9, 10, 11, 12],
    [13, 14, 15, 16]
  ],
  [1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 8, 11, 10]
)
