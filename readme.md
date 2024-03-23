# Gyan

A simple API to fetch description from Wikipedia and images from Google about {You say what}

<br />

## Usage
Please consider hosting this API on your own server.


`get /cat`
```json
    {
        "name": "cat",
        "link":"https://en.wikipedia.org/wiki/cat",
        "description":"The cat (Felis catus) is a domestic species of small carnivorous mammal. It is the only domesticated species in the family Felidae and is often referred to as the domestic cat to distinguish it from the...", 
        "images": [
            "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQtUjizJshX52X_qhGXsa2rIVOnGlRtXhBqnNPMB2JSn_ibAnT0s9aWUeAbfjE\u0026s",
            "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR7UDPBCebpUM7tMLJEJwrU_6DW6OBmLaLBVSmWM8YXq7ChWw2fPbR4ZrKEiA\u0026s",
            "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQDMskdkM6NS4oWAlHkAQ2giiArqn9evNFoTp2ExtsB321P8Db83BIT45cgfjI\u0026s",
            "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSLB8lwPF4UIM4Ck4ZdvI3qwYCX9InrEGdmqbcaQOac4CdOr4QZMDHWWFKjwNA\u0026s"
        ]
    }
```

<br />

## Issues
If you find any bugs/issues please report them by opening a new issue [here](https://github.com/sarthakpranesh/Gyan/issues)
