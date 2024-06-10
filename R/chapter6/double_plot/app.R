library("shiny")

library("ggplot2")

freqpoly <- function(x1, x2, binwidth = 0.1, xlim = c(-3, 3)) {
  df <- data.frame(
    x = c(x1, x2),
    g = c(rep("x1", length(x1)), rep("x2", length(x2)))
  )

  ggplot(df, aes(x, colour = g)) +
    geom_freqpoly(binwidth = binwidth, size = 1) +
    coord_cartesian(xlim = xlim)
}

ui <- fluidPage(
  fluidRow(
    column(6, plotOutput("hist")),
    column(6, plotOutput("hist2"))
  ),
  fluidRow(
    column(6, 
        numericInput("lambda1", label = "lambda1", value = 3),
        numericInput("lambda2", label = "lambda2", value = 5),
        numericInput("n", label = "n", value = 1e4, min = 0),
        actionButton("simulate", "Simulate!")
    ),
    column(6,
        numericInput("m", "Number of samples:", 2, min = 1, max = 100)
    )
  )
)

server <- function(input, output, session) {
    x1 <- eventReactive(input$simulate, {
        rpois(input$n, input$lambda1)
    }, ignoreNULL = FALSE)
    
    x2 <- eventReactive(input$simulate, {
        rpois(input$n, input$lambda2)
    }, ignoreNULL = FALSE)

    output$hist <- renderPlot({
        freqpoly(x1(), x2(), binwidth = 1, xlim = c(0, 40))
    }, res = 96)

    output$hist2 <- renderPlot({
        means <- replicate(1e4, mean(runif(input$m)))
        hist(means, breaks = 20)
    }, res = 96)

}


shinyApp(ui, server)