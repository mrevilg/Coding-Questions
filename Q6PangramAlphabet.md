# Alphabet Pangram Remainder

### Create an algorythm to detect not only if a word is a 'Pangram', but if not, what letters from the alp[habet are missing.

Assume ASCii, no whitespace, no special characters, result must be lower case.


```Csharp
namespace getMissingLetters
{
    class Program
    {
        static void Main(string[] args)
        {
            var userInput = "";

            Console.WriteLine("Please enter your sentence");
            userInput = Console.ReadLine();

            var modUserInput = (userInput.ToLower().Replace(" ", ""));

            String str = "abcdefghijklmnopqrstuvwxyz";
            String mask_str = modUserInput;
            Console.WriteLine(removeDirtyChars(str, mask_str));

        }

        static int NO_OF_CHARS = 250;

        static int[] getCharCountArray(String str)
        {
            int[] count = new int[NO_OF_CHARS];
            for (int i = 0; i < str.Length; i++) count[str[i]]++;
            return count;
        }

        static String removeDirtyChars(String str, String mask_str)
        {
            int[] count = getCharCountArray(mask_str);
            int ip_ind = 0, res_ind = 0;
            char[] arr = str.ToCharArray();
            while (ip_ind != arr.Length)
            {
                char temp = arr[ip_ind];
                if (count[temp] == 0)
                {
                    arr[res_ind] = arr[ip_ind];
                    res_ind++;
                }
                ip_ind++;
            }

            str = new String(arr);
            return str.Substring(0, res_ind);
        }
    }
}
```
| Input Text | Output |
| --- | --- |
| "A quick brown fox jumps over the lazy dog" | "" |
| "A slow yellow fox crawls under the proactive dog" | "bjkmqz" |
| "Lions, and tigers, and bears, oh my!" | "cfjkpquvwxz" |
| "" | "abcdefghijklmnopqrstuvwxyz" |
