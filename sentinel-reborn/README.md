# What the program does

    Our Program reads through the string of an input file, applies the modifiers in our code, formats it and writes the output file containing the transformed strings.


# How to run it with examples

    $ cat input.txt
    it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.
    $ go run . input.txt output.txt
    $ cat output.txt
    It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.
    $ cat input.txt
    Simply add 42 (hex) and 10 (bin) and you will see the result is 68.
    $ go run . input.txt output.txt
    $ cat output.txt
    Simply add 66 and 2 and you will see the result is 68.
    $ cat input.txt
    There is no greater agony than bearing a untold story inside you.
    $ go run . input.txt output.txt
    $ cat output.txt
    There is no greater agony than bearing an untold story inside you.
    $ cat input.txt
    Punctuation tests are ... kinda boring ,what do you think ?
    $ go run . input.txt output.txt
    $ cat output.txt
    Punctuation tests are... kinda boring, what do you think?

# Each transformation listed and explained

    Base

        Converts different bases (hexadecimal and binary) to decimal 

    article

        Fix articles where a changes to 'an' if a vowel is present at the beginning of the word after and remains 'a' if its a consonant

    quotes

        Fix spaces after and before quotes

    punctuation

        Joins all punctuations to the words before it without a space, and adds a space if there is a word after it.

    cases

        Converts words to up(converts assigned words to uppercase), low(converts assigned words to lowercase), and cap(converts first letter of every word to uppercase and the rest to lower)



# What your personal contribution was

    Coding Legend Alias Directer X: Eyung Emmanuel

        Director X was actively nudging and guiding us down the right path. He was patient from the planning phase to solo build and then the integration phase, down to assembly phase. In summary, he is the glue that held us together, without his insightful pointers, both we and our code crashes.

    Group Leader: Peace Oigoga

        Contributed to the Base converting function but most importantly, her subtle presence and underlying wisdom held us together through the storm. Not much of a talker but she spoke up in the moments that mattered especially the difficult ones.

    Assistant Group Leader: Treasure Gabriel

        A charismatic Develeoper in the making. She doubled down on the Base converting function alongside our amiable Leader. Also not much of talker, but like the leader that she is, spoke when it mattered and was quick to feast on the bugs we encountered.

    Chubiyojor Akoh

        A silent Bulldozer, with a brilliant mind that rivals Tesla. Tackled the Cases function without grumbling and he delivered a flawless code that runs without bugs. Not much of a talker but whenever he spoke we all stopped to marvel at his genuis.

    Chekwube Okafor

        A tireless workaholic who is calm, humble and patient. He built the punctuation function and fixed the bugs that came with it. A vital member of our squad that is integral to our functionality. 
    
    Faith Ejembi

        One of the most quiet members of the squad, she co-built the punctuation function and provided insight on the bugs that followed while actively testing for edge cases.

    Celebrate Owulo

        If small but mighty were to be a person, it would be Celebrate. A silent genuis with a broad knowledge of the basics of programming, it is safe to say she manipulates logic for a living. Built the  Article function and helped in making vital decisions that shaped our code.

    Alphonsus Ortse

        Another silent wheel that makes the Benchmarking machine run smoothly, also an avid Arsenal fan with no trophy. He co-built the Article function and helped tackle edge cases while we battled forking the main.

    Endurance Ochefije

        An integral member of the squad ever ready to take on tasks without grumbling, quiet and hardworking. She built the case function that ran without any bugs and is always happy to let the spotlight be on others while taking the backseat, a true Benchmarker through and through.

    Nzekwe Chinazaekpere

        An eloquent Benchmarker, ever ready to share her brilliant ideas. She understands the core basics of programming and she lives and breathes code. She built the Quote function and was quick to jump at any bug that popped and we almost felt sorry for the bugs. A proper Benchmarker.

    Emaikwu Godwin

        Co-built the Quote function, helped in fixing the branch issue an compilation of the README file. A jolly fellow always ready to share a joke and ease the tension.  


    
# One thing you found hardest today

    The thing we found hardest today was committing to a branch and merging the commits. As that took us over four hours as we had to resolve conflicts and begin anew almost five hours into the program. We admit that tensions were super high 🤒🥵

    If there was to be any point where we would have given up, it would have been here. But The Benchmarkers, ever ready to make it fly never quit, so we forged on and surmounted the hurdle. 

# One thing you understand now that you did not understand this morning

    The major things we picked up today that we didn't understand this morning are;creating a branch, committing to the branch and merging our commits to the main.
    We now know how to do all of these seamlessly although it was never an easy battle. 😞🥵