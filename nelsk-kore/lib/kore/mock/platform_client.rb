module Kore
  module Mock
    class PlatformClient
      attr_accessor :message_handler
      USER = 'irc_duder'

      def initialize(name)
        @name = name
      end

      def listen
        EM.defer do
          loop do
            input = gets.chomp
            message_handler.call("#{USER}: #{input}")
          end
        end
      end
    end

    def send(msg)
      puts "PlatformClient[#{self.name}] -> #{msg}"
    end
  end
end
