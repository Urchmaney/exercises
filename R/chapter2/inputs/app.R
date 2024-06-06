library(shiny)


datasets <- c("economics", "faithfuld", "seals")
ui <- fluidPage(
  selectInput("dataset", "Dataset", choices = list(Eastern = datasets, Western = c("Mike", "Lookas"))),
  textInput("iput", NULL, placeholder = "Your Name"),
  sliderInput("slid", "When should we deliver?", min = 0, max = 100, value = 10, step = 5, pre = "%F", animate = animationOptions()),
   sliderInput("slid", "When should we deliver?", min = as.Date("2024-01-01", "%Y-%m-%d"), max = as.Date("2024-12-12", "%Y-%m-%d"), value = as.Date("2024-05-05", "%Y-%m-%d"), timeFormat = "%F")

)

server <- function(input, output, session) {

}

shinyApp(ui, server)