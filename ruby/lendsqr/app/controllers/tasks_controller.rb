class TasksController < ApplicationController
  protect_from_forgery except: [:trello_web_hook]

  def index
    @tasks = helpers.fetch_lists_and_cards
  end

  def create
    helpers.create_new_trello_card(params[:name], params[:list_id])
    render 'index'
  end

  def new
    @list_id = params[:list_id]
  end

  def trello_web_hook
    if params["task"]['action']['display']['translationKey'] == 'action_move_card_from_list_to_list'
      puts '===================='
      puts ''
      puts ''
      puts "Card '#{params["task"]['action']['display']['entities']['card']['text']}' moved to '#{params["task"]['action']['display']['entities']['listAfter']['text']}'"
      puts ''
      puts ''
      puts '======================='
    end
    
    render json: { status: "Ok"}, status: :ok
  end
end