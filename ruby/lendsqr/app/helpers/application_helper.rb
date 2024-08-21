module ApplicationHelper
  def create_new_trello_card(name, list_id)
    url = "https://api.trello.com/1/cards?idList=#{list_id}&name=#{name}&key=#{ENV['TRELLO_KEY']}&token=#{ENV['TRELLO_TOKEN']}"
    response = RestClient.post(url, {})
  rescue StandardError => e
    p 'Error Creating card'
  end

  def fetch_lists_and_cards
    categorize_cards_to_list(fetch_board_lists, fetch_board_cards)
  end

  private

  def categorize_cards_to_list(lists, cards)
    list_hash = lists.map { |item| [item['id'], { id: item['id'], name: item['name'], cards: []}] }.to_h
    cards.each { |card| list_hash[card['idList']][:cards].push(card['name']) }
    list_hash.values
  end

  def fetch_board_cards
    url = "https://api.trello.com/1/boards/#{ENV['TRELLO_BOARD_ID']}/cards?key=#{ENV['TRELLO_KEY']}&token=#{ENV['TRELLO_TOKEN']}"
    response = RestClient.get(url)
    JSON.parse(response.body)
  rescue
    p 'Error Getting Cards'
  end

  def fetch_board_lists
    url = "https://api.trello.com/1/boards/#{ENV['TRELLO_BOARD_ID']}/lists?key=#{ENV['TRELLO_KEY']}&token=#{ENV['TRELLO_TOKEN']}"
    response = RestClient.get(url)
    JSON.parse(response.body)
  rescue
    p 'Error Getting List'
  end
end
