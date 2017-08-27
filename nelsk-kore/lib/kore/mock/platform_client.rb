module Kore
  module Mock
    class PlatformClient
      attr_accessor :message_handler

      def initialize(name)
        @name = name
        @user = "#{name}_duder"
      end

      def trigger(msg)
        self.message_handler.call(msg)
      end

      def send(msg)
        puts "PlatformClient[#{@name}] -> #{msg}"
      end
    end
  end
end
