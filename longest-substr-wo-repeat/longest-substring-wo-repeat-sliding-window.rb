class LongestSubstrSlidingWindow
    def length_of_longest_substring(s)
        char_map = {} # Stores: character => last_seen_index
        max_len = 0
        left = 0

        s.each_char.with_index do |char, right|
            # If we've seen this char and it's inside our current window
            if char_map.key?(char) && char_map[char] >= left
            left = char_map[char] + 1
            end

            # Update the last seen position
            char_map[char] = right
            
            # Calculate window size and update max
            current_window_len = right - left + 1
            max_len = current_window_len if current_window_len > max_len
        end

        max_len
    end
end
