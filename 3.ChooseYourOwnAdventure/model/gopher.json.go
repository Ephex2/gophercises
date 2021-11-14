package model

const DefaultAdventureString string = `
{
    "intro": {
      "title": "The Little Blue Gopher",
      "story": [
        "Once upon a time, long long ago, there was a little blue gopher. Our little blue friend wanted to go on an adventure, but he wasn't sure where to go. Will you go on an adventure with him?",
        "One of his friends once recommended going to New York to make friends at this mysterious thing called \"GothamGo\". It is supposed to be a big event with free swag and if there is one thing gophers love it is free trinkets. Unfortunately, the gopher once heard a campfire story about some bad fellas named the Sticky Bandits who also live in New York. In the stories these guys would rob toy stores and terrorize young boys, and it sounded pretty scary.",
        "On the other hand, he has always heard great things about Denver. Great ski slopes, a bad hockey team with cheap tickets, and he even heard they have a conference exclusively for gophers like himself. Maybe Denver would be a safer place to visit."
      ],
      "options": [
        {
          "text": "That story about the Sticky Bandits isn't real, it is from Home Alone 2! Let's head to New York.",
          "arc": "new-york"
        },
        {
          "text": "Gee, those bandits sound pretty real to me. Let's play it safe and try our luck in Denver.",
          "arc": "denver"
        }
      ]
    },
    "new-york": {
      "title": "Visiting New York",
      "story": [
        "Upon arriving in New York you and your furry travel buddy first attempt to hail a cab. Unfortunately nobody wants to give a ride to someone with a \"pet\". They kept saying something about shedding, as if gophers shed.",
        "Unwilling to accept defeat, you pull out your phone and request a ride using <undisclosed-app>. In a few short minutes a car pulls up and the driver helps you load your luggage. He doesn't seem thrilled about your travel companion but he doesn't say anything.",
        "The ride to your hotel is fairly uneventful, with the exception of the driver droning on and on about how he barely breaks even driving around the city and how tips are necessary to make a living. After a while it gets pretty old so you slip in some earbuds and listen to your music.",
        "After arriving at your hotel you check in and walk to the conference center where GothamGo is being held. The friendly man at the desk helped you get your badge and you hurry in to take a seat.",
        "As you head down the aisle you notice a strange man on stage with a mask, cape, and poorly drawn abs on his stomach. Next to him is a man in a... is that a fox outfit? What are these two doing? And what have you gotten yourself into?"
      ],
      "options": [
        {
          "text": "This is getting too weird for me. Let's bail and head back home.",
          "arc": "home"
        },
        {
          "text": "Maybe people just dress funny in the big city. Grab a a seat and see what happens.",
          "arc": "debate"
        }
      ]
    },
    "debate": {
      "title": "The Great Debate",
      "story": [
        "After a bit everyone settles down the two people on stage begin having a debate. You don't recall too many specifics, but for some reason you have a feeling you are supposed to pick sides."
      ],
      "options": [
        {
          "text": "Clearly that man in the fox outfit was the winner.",
          "arc": "sean-kelly"
        },
        {
          "text": "I don't think those fake abs would help much in a feat of strength, but our caped friend clearly won this bout. Let's go congratulate him.",
          "arc": "mark-bates"
        },
        {
          "text": "Slip out the back before anyone asks us to pick a side.",
          "arc": "home"
        }
      ]
    },
    "sean-kelly": {
      "title": "Exit Stage Left",
      "story": [
        "As you begin walking up to the fox-man you hear him introduce himself as Sean Kelly. While waiting in line you decide to do a little research to see what types of work Sean is into.",
        "A few clicks later and you drop your phone in horror. This guy's online handle is \"StabbyCutyou\". The stories about New York being dangerous were true!",
        "Without a thought you grab your gopher buddy and head for the door. \"I'll explain when we get to the hotel\" you tell him.",
        "After arriving at your hotel you both decide that you have had enough adventure. First thing tomorrow morning you are heading home."
      ],
      "options": [
        {
          "text": "You change your flight to leave early and head to the airport in the morning.",
          "arc": "home"
        }
      ]
    },
    "mark-bates": {
      "title": "Costume Time",
      "story": [
        "After talking with the wannabe superhero for a while you come to learn that his name is Mark Bates, and aside from his costume obsession he seems like a nice enough guy.",
        "It turns our Mark has been working on this project called Buffalo and he is desperately looking for a mascot. He even purchased a little buffalo outfit, but it is too small for him.",
        "After looking over the costume you are certain it won't fit you, but as luck would have it your gopher companion fit in it perfectly. Mark quickly snapped a few photos, mumbling something about Ashley McNamara designing the best buffalo costume ever.",
        "Many great times are had with Mark and pals, but you eventually find yourself on the last night of your stay."
      ],
      "options": [
        {
          "text": "Pack your bags and head to bed. We have a long flight in the morning.",
          "arc": "home"
        }
      ]
    },
    "denver": {
      "title": "Hockey and Ski Slopes",
      "story": [
        "You arrive in Denver and start your trip by attending a hockey game. The Avalanche had a rough season last year, but your gopher buddy is hopeful that they will do better this year. He also explains that he is tired of hearing about \"Two time Stanley Cup champion Phil Kessel.\" You suspect that he is still a little salty about the Penguins beating the San Jose Sharks in the Stanley Cup, but you decide to give him a break.",
        "The next day you head to the slopes and blaze a few trails. You can definitely see why Denver is called the \"Mile-High City\". Gorgeous mountain scenery doesn't come close to describing it.",
        "You consider checking out this GopherCon you have heard so much about, but a quick check on their website has your gopher buddy vetoing it. It turns out he has a strict, \"No Space Walks\" policy, and refuses to believe that it is just a graphic on the website.",
        "The week quickly flies by and before you know it you are packing up to head home."
      ],
      "options": [
        {
          "text": "Pack your bags and head to bed. We have a long flight in the morning.",
          "arc": "home"
        }
      ]
    },
    "home": {
      "title": "Home Sweet Home",
      "story": [
        "Your little gopher buddy thanks you for taking him on an adventure. Perhaps next year you can look into travelling abroad - you have both heard that gophers are all the rage in China."
      ],
      "options": []
    }
}
`
