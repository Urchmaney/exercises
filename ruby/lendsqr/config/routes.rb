Rails.application.routes.draw do
  # Define your application routes per the DSL in https://guides.rubyonrails.org/routing.html

  # Defines the root path route ("/")
  root "tasks#index"
  resources :tasks, only: [:create, :new] do
    collection do
      post 'trello_web_hook'
    end
  end
end
