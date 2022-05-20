# az-functions-custom-handlers-timer

    go build handler.go
    
    func start --verbose
    
Make sure you have a proper structure for your function app. Refer to this Microsoft doc to follow the steps, but choose "Timer Trigger" when first creating the function. The "local.settings.json" file is not in this repo, but will be created automatically for you if you follow the tutorial in the doc.

Basically, follow the doc until you have the function working locally. Replace the code in the "handler.go" file with the code in here.

https://docs.microsoft.com/en-us/azure/azure-functions/create-first-function-vs-code-other?tabs=go%2Cwindows
