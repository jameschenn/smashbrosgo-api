<h1 align="center">Smash Bros GO API</h1>
<p align="center"><img src="https://res.cloudinary.com/dttjhvohj/image/upload/v1717614247/smashbrosgo_banner.png" alt="Material Bread logo"></p>

Smash Bros GO API is a RESTful API built with everyone's favorite gopher language [GO](https://go.dev/doc/) aka Golang <img width=33px src="https://user-images.githubusercontent.com/25181517/192149581-88194d20-1a37-4be8-8801-5dc0017ffbbe.png">,
the [Gin](https://gin-gonic.com/docs/introduction/) Web Framework, [GORM](https://gorm.io/) ORM, and a PostgreSQL database. Using this API you can get information on fighters from Super Smash Bros. Ultimate. In the words of Masahiro Sakurai, **EVERYONE IS HERE** 

## Getting Started

1. Clone the Project:

  ```bash
  git clone https://github.com/jameschenn/smashbrosgo-api.git
  ```
2. Download Dependencies:

  ```bash
  go mod tidy
  ```
3. Create Postgresql database. I personally used and love using [Supabase](https://supabase.com/)

4. Create .env file with **PORT** and **DATABASE** variables
  ```bash
  PORT=3000
  DATABASE={YOUR DATABASE CREDENTIALS}
  ```
5. Migrate tables to database
  
  ```bash
  go run migrations/migrations.go
  ```

6. Seed Database
  
  ```bash
  go run seeders/seed.go
  ```

6. Run the API:
  > [!NOTE]
  > [CompileDaemon](https://pkg.go.dev/github.com/githubnemo/CompileDaemon#section-readme) watches your .go files in a directory and invokes go build if a file changed. 

  ```bash
  CompileDaemon -command="./smashbrosgo-api"
  ```

## Routes 

**GET** Fetch All Characters
  ```bash
  http://localhost:3000/characters
  ```

**GET** Fetch Character by ID
  ```bash
  http://localhost:3000/characters/{ID}
  ```

**POST** Create Character
  ```bash
  http://localhost:3000/characters
  ```
  ```
    {
        "Name": "James Chen",
        "Series": "James Chen",
        "FirstAppearance": "Earth",
        "Description": "test",
        "Tier": "S",
        "Weight": 160,
        "Speed": 3.33,
        "NeutralB": "Code",
        "UpB": "Rising Collaboration",
        "SideB": "Strong Communication",
        "DownB": "Critical Thinking Charge",
        "Final_Smash": "Got the Job!!",
        "Image_URL": "https://avatars.githubusercontent.com/u/73676915?v=4"
    }
  ```

**UPDATE** Edit Character by ID
  ```bash
  http://localhost:3000/characters/{ID}
  ```
  ```
  {
	  "Tier": "S+"
  }
  ```

**DELETE** Delete Character by ID
  ```bash
  http://localhost:3000/characters/{ID}
  ```
