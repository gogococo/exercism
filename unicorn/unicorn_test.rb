gem 'minitest', '~> 5.0'

require 'minitest/autorun'
require 'minitest/pride'

require_relative 'unicorn'

class UnicornTest < Minitest::Test
  def test_it_has_a_name
    unicorn = Unicorn.new("Robert")
    # assert_equal expected_value, actual_value
    assert_equal "Robert", unicorn.name
  end

  def test_it_can_be_named_something_else
    unicorn = Unicorn.new("Joe")
    assert_equal "Joe", unicorn.name
  end

  def test_it_is_white_by_default
    unicorn = Unicorn.new("Margie")
    assert_equal "white", unicorn.color
  end

  def test_it_knows_if_it_is_white
    unicorn = Unicorn.new("Liz")
    #assert_equal true, unicorn.white?
    assert unicorn.white?, "Liz should be white, but isn't"
  end

  def test_it_does_not_have_to_be_white
    unicorn = Unicorn.new("Barb", "purple")
    assert_equal "purple", unicorn.color
  end

  def test_it_knows_if_it_is_white
    unicorn = Unicorn.new("Roxy", "green")
    #assert_equal false, unicorn.white?
    refute unicorn.white?, "I guess Roxy thinks she's white, when really she is green."
  end

  def test_unicorn_says_sparkly_stuff
    unicorn = Unicorn.new("Johnny")
    assert_equal "**;* Wonderful! **;*", unicorn.say("Wonderful!")
  end

  def test_unicorn_says_different_sparkly_stuff
    unicorn = Unicorn.new("Francis")
    assert_equal "**;* I don't like you very much. **;*", unicorn.say("I don't like you very much.")
  end

end
