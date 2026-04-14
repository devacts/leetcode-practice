class LongestSubstr
    def self.length_of_longest_substring(s)
        curr_substr = ''
        longest_sub_len = 0
        s.each_char do |c|
            unless curr_substr.include? c
                curr_substr.concat(c)
            else
                if curr_substr.length > longest_sub_len
                    longest_sub_len = curr_substr.length
                end
                index = curr_substr.index(c)
                if curr_substr[index + 1]
                    curr_substr = curr_substr[index + 1..].concat(c)
                else
                    curr_substr = c
                end
            end
        end
        [longest_sub_len, curr_substr.length].max
    end
end

#s = 'abcabcbb'
#s = 'abcc'
#s = 'dvdf'
s = ' '
#s = 'pwwkew'
p LongestSubstr.length_of_longest_substring(s)
