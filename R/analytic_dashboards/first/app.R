library("shiny")

ui <- fluidPage(
    fluidRow(
        column(2),
        column(10,
            fluidPage(
                titlePanel("Hello Shiny!"),
                sidebarLayout(
                    sidebarPanel(
                    sliderInput("obs", "Observations:", min = 0, max = 1000, value = 500)
                    ),
                    mainPanel(
                        div(
                            p("Text")
                        )
                    plotOutput("distPlot")
                    )
                )
            )
        )
    )
) 


server <- function(input, output, session) {

}

shinyApp(ui, server)