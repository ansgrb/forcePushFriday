# forcePushFriday meme generator
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

*because who needs a peaceful weekend?* 😈

this go program generates dark humor ascii art memes that encourage the chaos of deploying to production on friday. perfect for keeping your coworkers on their toes and your devops team in a constant state of anxiety.

## 👀 use cases

### 🎲 basic random generation

```bash
$ ./forcePushFriday
╔════════════════════════════════════════════════╗
║     deploy on friday, debug on monday          ║
╚════════════════════════════════════════════════╝
        (•̀ᴗ•́)و       
         /|\         
        / | \        
         / \         
        /   \        
       👟    👟       

  • production fire risk: 87%
  • angry manager likelihood: 94%
  • coffee cups required for aftermath: 7
```

### 💬 custom message with specific style

```bash
$ ./forcePushFriday -message "yolo deploy friday" -style 0 -figure 3
┌──────────────────────────────────────────────┐
│              yolo deploy friday              │
└──────────────────────────────────────────────┘
      _____           
     /     \          
    | ^   ^ |         
    |   ω   |         
     \_____/          
        ||            
       /||\           
      / || \          
        ⅃⅃ 

  • career impact rating: 9/10
  • monday morning regret factor: 8/10
```

### ⏰ friday afternoon detection

```bash
$ ./forcePushFriday -width 60
*************************************************************
*     friday deployments: the russian roulette of software   *
*************************************************************
            _\/_        
             /\         
             /\         
            /  \        
           |    |       
           |    |       
         __|    |__     
        /__|    |__\    

  • weekend ruined probability: 92%
  • stack overflow visits needed: 8

  ⚠️  it's actually friday afternoon right now. do it. do it.
```

## ✨ features

- 📝 **random message templates**: choose from 14+ pre-defined dark humor messages about friday deployments or create your own
- 🎨 **customizable ascii art**: 5 different box styles and 5 character figures
- 📊 **chaos metrics**: randomly generated statistics about the potential fallout of your friday deploy
- 🌈 **background patterns**: add visual flair with optional background patterns
- 🎭 **color support**: colorful output to maximize attention and minimize regret
- 🕒 **friday detection**: extra encouragement if it's actually friday afternoon when you run it

## 🛠️ installation

```bash
# clone the repository
git clone https://github.com/ansgrb/forcePushFriday.git
cd forcePushFriday

# install dependencies
go get github.com/fatih/color

# build the program
go build forcePushFriday

# optional: move to your path
mv forcePushFriday /usr/local/bin/
```

## 🚦 usage

### 🏁 basic usage

```bash
# generate a random meme
./forcePushFriday

# use a custom message
./forcePushFriday -message "yolo deploy friday"
```

### 🎛️ command line options

| flag            | description                             | default               |
|-----------------|-----------------------------------------|-----------------------|
| `-message`      | custom message for the meme             | random from templates |
| `-color`        | use colorful output                     | `true`                |
| `-style`        | box style (0-4, -1 for random)          | `-1` (random)         |
| `-figure`       | figure style (0-4, -1 for random)       | `-1` (random)         |
| `-bg`           | background pattern (0-7, -1 for random) | `-1` (random)         |
| `-chaos`        | chaos level data (0-5, -1 for random)   | `-1` (random)         |
| `-width`        | width of the meme box                   | `50`                  |
| `-list-figures` | display all available figures           | `false`               |
| `-help`         | display help information                | `false`               |
### 🧙‍♂️ advanced examples

```bash
# full customization with lots of chaos
./forcePushFriday -message "breaking prod like a boss" -style 2 -figure 4 -bg 4 -chaos 5 -width 60

# no colors (for terminal purists)
./forcePushFriday -color=false

# display help and options
./forcePushFriday -help
```

## 💌 using in team communications

share memes with your team directly from the terminal:

```bash
# send to slack via webhook
./forcePushFriday | curl -X POST -H "Content-type: text/plain" --data-binary @- https://hooks.slack.com/services/YOUR/SLACK/WEBHOOK

# email to the team before leaving for the weekend
./forcePushFriday | mail -s "just pushed to production!" team@example.com

# add to your deployment scripts
deploy_to_production && ./friday-deploy || echo "deployment failed but at least it's friday!"
```

## ✏️ adding custom messages

to add your own message templates, modify the `messageTemplates` slice in the source code:

```go
var messageTemplates = []string{
    // existing messages...
    "your custom message here",
    "another custom message",
}
```

## 🔍 implementation details

the program uses several components to generate the memes:

1. 📋 **message selection**: either uses your custom message or randomly selects from pre-defined templates
2. 📦 **box style generation**: creates a box around the message using ascii/unicode characters
3. 👤 **figure selection**: adds a cheerful ascii character below the message
4. 📈 **chaos metrics**: generates random metrics about the consequences of friday deployments
5. 🎭 **background patterns**: optional patterns to make the meme stand out

## 🤝 contributing

contributions are welcome! please feel free to submit a pull request.

1. fork the repository
2. create your feature branch (`git checkout -b feature/amazing-feature`)
3. commit your changes (`git commit -m 'add some amazing feature'`)
4. push to the branch (`git push origin feature/amazing-feature`)
5. open a pull request

## ☃️ available figures (add yours)
```bash

Figure 0:
          _\/_        
           /\         
           /\         
          /  \        
         |    |       
         |    |       
       __|    |__     
      /__|    |__\    

Figure 1:
        (•̀ᴗ•́)و       
         /|\         
        / | \        
         / \         
        /   \        
       👟    👟       

Figure 2:
       \(^o^)/       
         |__|         
          ||          
          ||          
         /  \         
        /    \        

Figure 3:
      _____           
     /     \          
    | ^   ^ |         
    |   ω   |         
     \_____/          
        ||            
       /||\           
      / || \          
        ⅃⅃            

Figure 4:
        (҂◡_◡)        
       ᕦ(╭ರ ⊙ ͜ʖ⊙ര)ᕥ    
          |           
         /|\          
        / | \         
         J L          

Figure 5:
       (╯°□°)╯        
       ┻━━┻           
        /|\           
       / | \          
        / \           
       /   \          

Figure 6:
        _ノ乙            
       ( ͡° ͜ʖ ͡°)       
       /  ⌒\         
      /   |  \        
     /   /|   \       
         / \          
        /   \         

Figure 7:
        ______        
       /      \       
      | ⊙  ⊙  |       
      |   ▽    |      
       \______/       
         |  |         
        ┌|  |┐        
        └|  |┘        
         |__|         
         /  \         

Figure 8:
        (ง •̀_•́)ง      
          |           
          |           
         /|\          

Figure 9:
       \( ⌐■_■)/      
         \| |/        
          | |         
         /| |\        
        / | | \       
          ⅃ Ⅎ          

Figure 10:
         /|\          
        /*|*\         
       /* | *\        
      /*  |  *\       
     /*   |   *\      
    /*    |    *\     
   /*************\    
          |  |         
          |  |         
         /|  |\       

Figure 11:
      ┌─┐             
      ┴─┴             
   ಠ_ಠ                
   <|>                
   /ω\                

Figure 12:
        _______       
       /       \      
      | ⇀ ‿ ↼  |      
      |  SOON   |     
       \_______/      
       W  | |  W      
          | |         
          | |         
         // \\        

Figure 13:
       .---------.    
      / .-------. \   
     / /         \ \  
     | |         | |  
    _| |_________| |_ 
  .' |_|         |_| '.
  '._____ ____ _____.'
  |     .'____'.     |
  '.__.'.'    '.'.__.'
  '.__  |      |  __.'
  |   '.'.____.'.'   |
  '.____'.____.'____.'
  '.________________.'

Figure 14:
         🔥  🔥         
        🔥    🔥        
       🔥 PROD 🔥       
        🔥    🔥        
         🔥  🔥         


```

## ⚠️ disclaimer

this tool is for humor purposes only. the author takes no responsibility for any actual friday deployments made under its influence. remember: with great power comes great deniability.

---

## 📜 License

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

---
### *"friday deploys are a sacrifice. monday is when the gods collect."*  🧙‍♂️